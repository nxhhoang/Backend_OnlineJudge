package utils

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"problem/models"
	"regexp"
	"strconv"
	"time"

	"github.com/xyproto/unzip"
)

/*
polygonApiCall(): make Polygon API calls
- Remember to do .Body.Close() the response
*/
func polygonApiCall(method string, params map[string]string) (*http.Response, error) {
	apiSecret := os.Getenv("POLYGON_API_SECRET")

	apiSig := ""

	// Create apiSig
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand_header := make([]byte, 6)
	for i := range 6 {
		rand_header[i] = charset[rand.Intn(len(charset))]
	}
	apiSig += fmt.Sprintf("%s/problem.package?", rand_header)

	keys := GetSortedKeys(&params, func(a, b string) bool { return a < b })
	for k, v := range keys {
		if k > 0 {
			apiSig += "&"
		}
		apiSig += fmt.Sprintf("%s=%s", v, params[v])
	}
	apiSig += fmt.Sprintf("#%s", apiSecret)

	sha512 := sha512.New()
	sha512.Write([]byte(apiSig))
	apiSig = hex.EncodeToString(sha512.Sum(nil))

	apiSig = string(rand_header) + apiSig

	address := fmt.Sprintf("https://polygon.codeforces.com/api/%s?", method)
	requestParams := url.Values{}
	requestParams.Add("apiSig", apiSig)
	for key, value := range params {
		requestParams.Add(key, value)
	}

	address += requestParams.Encode()

	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
DownloadPackge() - Download Polygon problems
- Problems are stored at $STORAGE_DIR/$problemId/ with structure following the README.md
- It removes all downloaded packages of the given problem before downloading the specified package

FUTURE:
*/
func DownloadPackage(problemId uint64, packageId uint64) error {
	params := map[string]string{
		"problemId": strconv.Itoa(int(problemId)),
		"packageId": strconv.Itoa(int(packageId)),
		"apiKey":    os.Getenv("POLYGON_API_KEY"),
		"type":      "linux",
		"time":      fmt.Sprintf("%d", time.Now().Unix()),
	}
	resp, err := polygonApiCall("problem.package", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}

	// Extract to a temporary directory
	f, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(body); err != nil {
		return err
	}

	dirpath := fmt.Sprintf("%s/%s", os.Getenv("PROBLEM_STORAGE_DIR"), params["problemId"])
	if err := os.RemoveAll(dirpath); err != nil {
		return err
	}
	if err := os.Mkdir(dirpath, os.ModePerm); err != nil {
		return nil
	}

	tempdir, err := os.MkdirTemp("", "")
	if err != nil {
		return err
	}
	defer os.Remove(tempdir)

	if err = unzip.Extract(f.Name(), tempdir); err != nil {
		return err
	}

	var xml *os.File
	if xml, err = os.Open(tempdir + "/problem.xml"); err != nil {
		return err
	}
	defer xml.Close()

	var problem models.Problem
	if problem, err = ParseProblemStruct(problemId, xml); err != nil {
		return err
	}

	if err := SaveProblemToJson(problem, dirpath+"/problem.json"); err != nil {
		return err
	}

	if err := os.Mkdir(dirpath+"/tests", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir(dirpath+"/tests/input", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir(dirpath+"/tests/output", os.ModePerm); err != nil {
		return err
	}

	r, err := regexp.Compile(`\.a$`)
	if err != nil {
		return nil
	}

	err = filepath.Walk(tempdir+"/tests/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		var destFile io.Writer

		if !r.Match([]byte(info.Name())) {
			destFile, err = os.Create(dirpath + "/tests/input/" + info.Name())
		} else {
			destFile, err = os.Create(dirpath + "/tests/output/" + info.Name()[:len(info.Name())-2]) // remove the .a extension
		}

		if err != nil {
			return err
		}

		if _, err := io.Copy(destFile, sourceFile); err != nil {
			return err
		}

		return nil
	})

	var errBuffer bytes.Buffer

	cmd := exec.Command("scripts/gen_statement/main.sh", tempdir, dirpath)
	cmd.Stderr = &errBuffer
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error creating statement: %s", errBuffer)
	}

	cmd = exec.Command("scripts/compile_checker/main.sh", tempdir, dirpath)
	cmd.Stderr = &errBuffer
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error compiling checker: %s", errBuffer)
	}

	return err
}
