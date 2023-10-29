package main

import (
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zcubbs/mq-watch/cmd/server/api"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/mqttclient"
	"gorm.io/gorm"
	"log"
)

var (
	configPath = flag.String("config", ".", "Path to the configuration file")
)

////go:embed web/views
//var viewFiles embed.FS
//
////go:embed web/assets
//var staticFiles embed.FS

func main() {
	flag.Parse()

	cfg, err := config.LoadConfiguration(*configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	config.PrintConfiguration(cfg)

	conn, err := db.InitializeDB(cfg.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	mqc, err := mqttclient.ConnectAndSubscribe(
		cfg.MQTT.Broker,
		cfg.MQTT.Topic,
		func(client mqtt.Client, msg mqtt.Message) {
			messageHandler(conn, msg)
		},
	)
	if err != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", err)
	}

	defer mqc.Disconnect(250)

	// init template engine
	//engine := html.NewFileSystem(http.FS(viewFiles), ".html")
	//engine.Engine.Directory = "web/views"
	//engine.Reload(false)

	// init server
	app := fiber.New(fiber.Config{
		//Views:                 engine,
		//ViewsLayout:           "layouts/main",
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//app.Use("/assets", filesystem.New(
	//	filesystem.Config{
	//		Root:       http.FS(staticFiles),
	//		PathPrefix: "web/assets",
	//		Browse:     false,
	//	},
	//))

	// API endpoint
	app.Get("/api/messages", func(c *fiber.Ctx) error {
		// You might need to adjust this part based on the api.MessageHandler function
		return api.MessageHandler(conn, c)
	})

	//// serve index
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.Render("index", fiber.Map{
	//		"Title": "MQ Watch",
	//	})
	//})

	// Run the server
	err = app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
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
		log.Printf("Error decoding message: %v", err)
		return
	}

	db.SaveMessage(conn, payload.Tenant, 1, payload.CreatedAt)
}
