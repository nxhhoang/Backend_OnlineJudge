package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bibimoni/Online-judge/gateway/src/common"
	"github.com/bibimoni/Online-judge/gateway/src/infrastructure/config"
)

type AuthResponseBody struct {
	Code int      `json:"code,omitempty"`
	User UserType `json:"user"`
}

type UserType struct {
	Username string `json:"username,omitempty"`
	Id       int    `json:"id,omitempty"`
}

func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authURL := config.Load().Endpoints.Auth

		token := extractBearer(r.Header.Get("Authorization"))
		res, err := common.SendRequest[AuthResponseBody](r.Context(), common.APIRequest{
			Method:  "GET",
			URL:     fmt.Sprintf("%s/auth/validate/%s", authURL, token),
			Timeout: 10 * time.Second,
		})

		if err != nil {
			if res == nil {
				http.Error(w, "An error occured", http.StatusInternalServerError)
				return
			}
			http.Error(w, "An error occured", res.StatusCode)
			return
		}

		// Add username field into body
		org, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var obj map[string]any
		if len(org) > 0 {
			if err := json.Unmarshal(org, &obj); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
			}
		} else {
			obj = make(map[string]any)
		}
		obj["username"] = res.PayLoad.User.Username

		newBytes, _ := json.Marshal(obj)

		r.Body = io.NopCloser(bytes.NewReader(newBytes))

		r.ContentLength = int64(len(newBytes))

		r.Header.Set("Content-Length", strconv.Itoa(len(newBytes)))

		// serve
		next.ServeHTTP(w, r)
	})
}

func extractBearer(header string) string {
	const prefix = "Bearer "
	return strings.TrimPrefix(header, prefix)
}
