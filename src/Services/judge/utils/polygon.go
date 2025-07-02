package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/xyproto/unzip"
)

/*
DownloadPackge() - Download Polygon problems
- Problems are stored at /upload/problems/$problemId/$packageId

FUTURE:
- Automatically get the latest packageId
*/
func DownloadPackage(params map[string]string) error {
	dirpath := fmt.Sprintf("/problems/%s/%s", params["problemId"], params["packageId"])
	if err := os.Mkdir(dirpath, 0755); os.IsExist(err) {
		return errors.New("the problem already existed")
	}

	apiSecret := os.Getenv("POLYGON_API_SECRET")
	if apiSecret == "" {
		return errors.New("POLYGON_API_SECRET not set")
	}

	// get 6 random bytes
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand_header := make([]byte, 6)
	for i := range 6 {
		rand_header[i] = charset[rand.Intn(len(charset))]
	}
	apiSig := fmt.Sprintf("%s/problem.package?", rand_header)

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

	// fmt.Printf("apiSig: %s\n", apiSig)

	// Build address

	address := "https://polygon.codeforces.com/api/problem.package?"
	requestParams := url.Values{}
	requestParams.Add("problemId", params["problemId"])
	requestParams.Add("packageId", params["packageId"])
	requestParams.Add("type", params["type"])
	requestParams.Add("apiKey", params["apiKey"])
	requestParams.Add("time", params["time"])
	requestParams.Add("apiSig", apiSig)

	address += requestParams.Encode()
	fmt.Printf("address: %s\n", address)

	// make requests
	resp, err := http.Get(address)
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

	f, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	// TODO: save it to dir path
	err = unzip.Extract(f.Name(), dirpath)
	if err != nil {
		return err
	}

	return nil
}
