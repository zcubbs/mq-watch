package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/zcubbs/mq-watch/cmd/server/api"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/mqttclient"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	configPath = flag.String("config", ".", "Path to the configuration file")
)

//go:embed web/index.html web/app.js
var content embed.FS

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

	// Setup and run the HTTP server
	r := gin.Default()
	r.GET("/api/messages", func(c *gin.Context) {
		api.MessageHandler(conn, c)
	})

	// Serve the UI files
	// Serve the static files
	r.GET("/web/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		data, err := content.ReadFile("web" + filepath)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/plain", data)
	})

	// Serve index.html at the root
	r.GET("/", func(c *gin.Context) {
		data, err := content.ReadFile("web/index.html")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/html", data)
	})

	err = r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

type MessagePayload struct {
	Device string `json:"device"`
	Tenant string `json:"tenant"`
	Data   string `json:"data"`
}

func messageHandler(conn *gorm.DB, msg mqtt.Message) {
	var payload MessagePayload
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Printf("Error decoding message: %v", err)
		return
	}

	db.SaveMessage(conn, payload.Tenant, 1)
}
