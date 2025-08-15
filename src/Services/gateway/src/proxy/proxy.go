package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/bibimoni/Online-judge/gateway/src/infrastructure/config"
)

func SubmissionApiProxy() http.Handler {
	cfg := config.Load()
	return newProxy(cfg.Endpoints.Submission)
}

func LoginApiProxy() http.Handler {
	cfg := config.Load()

	return newProxy(cfg.Endpoints.Auth)
}

func newProxy(endpoint string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, err := url.Parse(endpoint)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		httpProxy := httputil.NewSingleHostReverseProxy(url)
		httpProxy.ServeHTTP(w, r)
	})
}
