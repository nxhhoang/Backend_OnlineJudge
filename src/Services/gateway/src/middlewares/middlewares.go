package middlewares

import (
	"net/http"

	"github.com/bibimoni/Online-judge/gateway/src/infrastructure/config"
)

func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		authURL := config.Load().Endpoints.Auth
	})
}
