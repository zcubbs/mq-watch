package main

import (
	"flag"
	"github.com/zcubbs/mq-watch/cmd/server/api"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"github.com/zcubbs/mq-watch/cmd/server/mqttclient"
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

	// init db store
	store := db.NewPostgresStore(conn)

	log.Info("Initializing MQTT client")
	mqc, err := mqttclient.ConnectAndSubscribe(cfg.MQTT, cfg.Tenants, store)
	if err != nil {
		log.Fatal("Error connecting to MQTT broker", "error", err)
	}

	defer mqc.Disconnect(250)

	// Start the server
	log.Info("Starting server")
	server := api.NewServer(cfg.Server, store)
	server.Start()
}
