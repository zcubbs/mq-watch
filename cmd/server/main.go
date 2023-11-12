package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/zcubbs/mq-watch/cmd/server/api"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"github.com/zcubbs/mq-watch/cmd/server/mqttclient"
	"gorm.io/gorm"
	"net/http"
)

var (
	configPath = flag.String("config", ".", "Path to the configuration file")
)

//go:embed web/dist/*
var webDist embed.FS

var (
	log = logger.L()
)

func main() {
	flag.Parse()

	cfg, err := config.LoadConfiguration(*configPath)
	if err != nil {
		log.Fatal("Error loading configuration", "error", err)
	}

	config.PrintConfiguration(cfg)

	log.Info("Connecting to database", "datasource", cfg.Database.Datasource)
	conn, err := db.InitializeDB(cfg.Database)
	if err != nil {
		log.Fatal("Error initializing database", "error", err)
	}

	mqc, err := mqttclient.ConnectAndSubscribe(
		cfg.MQTT,
		cfg.Tenants,
		func(client mqtt.Client, msg mqtt.Message) {
			messageHandler(conn, msg)
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
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Serve the web app
	app.Use("/", filesystem.New(
		filesystem.Config{
			Root:       http.FS(webDist),
			PathPrefix: "web/dist",
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

	// Run the server
	log.Info("Starting server", "port", cfg.Server.Port)
	err = app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatal("Error starting server", "error", err)
	}
}

type MessagePayload struct {
	Device    string `json:"device"`
	Tenant    string `json:"tenant"`
	CreatedAt string `json:"created_at"`
	Data      string `json:"data"`
}

func messageHandler(conn *gorm.DB, msg mqtt.Message) {
	var payload MessagePayload
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Error("Error decoding message", "error", err)
		return
	}

	log.Info("Received message", "payload", payload)

	db.SaveMessage(conn, payload.Tenant, 1, payload.CreatedAt)
}
