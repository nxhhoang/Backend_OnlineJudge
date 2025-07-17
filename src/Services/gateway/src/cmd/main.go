package main

import (
	"github.com/bibimoni/Online-judge/gateway/src/infrastructure/config"
	"github.com/bibimoni/Online-judge/gateway/src/proxy"
	"github.com/bibimoni/Online-judge/gateway/src/server"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	config.NewLogger(cfg.LogLevel)
	s := server.NewServer()
	r := server.GetRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", proxy.SubmissionApiProxy())
	s.ListenAndServe()
}
