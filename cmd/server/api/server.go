package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"github.com/zcubbs/mq-watch/cmd/server/web"
	"net/http"
)

var log = logger.L()

type Server struct {
	store db.Store
	app   *fiber.App
	cfg   config.ServerConfiguration
}

func NewServer(cfg config.ServerConfiguration, store db.Store) *Server {
	return &Server{
		store: store,
		app:   initApp(),
		cfg:   cfg,
	}
}

func initApp() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	return app
}

func (s *Server) Start() {
	s.mountApiRoutes()
	s.mountStaticRoutes()

	// Run the server
	log.Info("Starting server", "port", s.cfg.Port)

	port := fmt.Sprintf(":%d", s.cfg.Port)

	if s.cfg.TlsEnabled {
		log.Info("Starting server with TLS enabled")
		log.Fatal("failed to run tls secure server",
			s.app.ListenTLS(port, s.cfg.TlsCertFile, s.cfg.TlsKeyFile))
	} else {
		log.Info("Starting server with TLS disabled")
		log.Fatal("failed to run server", s.app.Listen(port))
	}
}

func (s *Server) mountApiRoutes() {
	s.app.Get("/health", healthHandler)
	s.app.Get("/api/messages", func(c *fiber.Ctx) error {
		return messageHandler(s.store, c)
	})
	s.app.Get("/api/total-messages-per-day", func(c *fiber.Ctx) error {
		return totalMessagesPerDayHandler(s.store, c)
	})
	s.app.Get("/api/top-tenants", func(c *fiber.Ctx) error {
		return getTopTenantsHandler(s.store, c)
	})
	s.app.Get("/api/message-stats", func(c *fiber.Ctx) error {
		return getMessageStatsHandler(s.store, c)
	})
	s.app.Post("/api/save-messages", func(c *fiber.Ctx) error {
		return saveMessagesHandler(s.store, c)
	})
}

func (s *Server) mountStaticRoutes() {
	s.app.Use("/*", filesystem.New(
		filesystem.Config{
			Root:       http.FS(web.SpaFiles),
			PathPrefix: "dist",
			Browse:     false,
		},
	))
}

func (s *Server) GracefulShutdown() {
	log.Info("Shutting down server")
	if err := s.app.Shutdown(); err != nil {
		log.Error("Error shutting down server", "error", err)
	}
}
