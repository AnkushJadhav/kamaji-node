package http

import (
	"time"

	"github.com/gofiber/fiber"
)

func (srv *Server) initServerSettings() {
	srv.settings = &fiber.Settings{}
}

func (srv *Server) prepopulatePool(b bool) {
	srv.settings.Prefork = b
}

func (srv *Server) setReadTimeout(t time.Duration) {
	srv.settings.ReadTimeout = t
}

func (srv *Server) setWriteTimeout(t time.Duration) {
	srv.settings.WriteTimeout = t
}

func (srv *Server) setMaxConcurrency(c int) {
	srv.settings.Concurrency = c
}

func (srv *Server) initServer() {
	srv.app = fiber.New(srv.settings)
}
