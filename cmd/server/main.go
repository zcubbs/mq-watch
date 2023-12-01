package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/zcubbs/mq-watch/cmd/server/api"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"github.com/zcubbs/mq-watch/cmd/server/mqttclient"
	"github.com/zcubbs/mq-watch/cmd/server/web"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var (
	configPath = flag.String("config", ".", "Path to the configuration file")
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var (
	log = logger.L()
)

func main() {
	flag.Parse()

	log.Info("Starting mq-watch", "version", Version, "commit", Commit, "date", Date)

	cfg, err := config.LoadConfiguration(*configPath)
	if err != nil {
		log.Fatal("Error loading configuration", "error", err)
	}

	config.PrintConfiguration(cfg)

	conn, err := db.InitializeDB(cfg.Database)
	if err != nil {
		log.Fatal("Error initializing database", "error", err)
	}

	log.Info("Initializing MQTT client")
	mqc, err := mqttclient.ConnectAndSubscribe(
		cfg.MQTT,
		cfg.Tenants,
		func(tMsg mqttclient.TenantMessage) {
			messageHandler(conn, tMsg)
		},
	)
	if err != nil {
		log.Fatal("Error connecting to MQTT broker", "error", err)
	}

	defer mqc.Disconnect(250)

	// init server
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Serve the web app
	app.Use("/", filesystem.New(
		filesystem.Config{
			Root:       http.FS(web.SpaFiles),
			PathPrefix: "dist",
			Browse:     false,
		},
	))

	// serve the api routes
	app.Get("/api/messages", func(c *fiber.Ctx) error {
		// You might need to adjust this part based on the api.MessageHandler function
		return api.MessageHandler(conn, c)
	})
	app.Get("/api/total-messages-per-day", func(c *fiber.Ctx) error {
		return api.TotalMessagesPerDayHandler(conn, c)
	})
	app.Get("/api/top-tenants", func(c *fiber.Ctx) error {
		return api.GetTopTenantsHandler(conn, c)
	})
	app.Get("/api/message-stats", func(c *fiber.Ctx) error {
		return api.GetMessageStatsHandler(conn, c)
	})
	app.Post("/api/save-messages", func(c *fiber.Ctx) error {
		return api.SaveMessagesHandler(conn, c)
	})

	// Run the server
	log.Info("Starting server", "port", cfg.Server.Port)

	port := fmt.Sprintf(":%d", cfg.Server.Port)

	if cfg.Server.TlsEnabled {
		log.Info("Starting server with TLS enabled")
		log.Fatal("failed to run tls secure server",
			app.ListenTLS(port, cfg.Server.TlsCertFile, cfg.Server.TlsKeyFile))
	} else {
		log.Info("Starting server with TLS disabled")
		log.Fatal("failed to run server", app.Listen(port))
	}
}

func messageHandler(conn *gorm.DB, msg mqttclient.TenantMessage) {
	log.Debug("Received message",
		"tenant", msg.Tenant,
		"topic", msg.Message.Topic(),
	)

	var payload string
	if msg.SavePayload {
		payload = string(msg.Message.Payload())
	} else {
		payload = ""
	}

	err := db.SaveMessage(
		conn, msg.Tenant,
		msg.Message.Topic(),
		payload,
		time.Now().Format(time.RFC3339),
	)

	if err != nil {
		log.Error("Error saving message",
			"tenant", msg.Tenant,
			"topic", msg.Message.Topic(),
			"error", err)
	}
}
