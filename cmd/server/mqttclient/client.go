package mqttclient

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
	"os"
)

var (
	log = logger.L()
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Debug("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info("Connected to MQTT Broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Error("Connect lost", "error", err)
}

// TenantMessage wraps an MQTT message with tenant information
type TenantMessage struct {
	Message     mqtt.Message
	SavePayload bool
	Tenant      string
}

func ConnectAndSubscribe(cfg config.MQTTConfiguration, tenants []config.TenantConfiguration, store db.Store) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	if cfg.TlsEnabled {
		opts.SetTLSConfig(newTLSConfig(cfg))
	}
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	for _, tenant := range tenants {
		for _, topic := range tenant.Topics {
			// Wrap the messageHandler to include tenant information
			wrappedHandler := func(client mqtt.Client, msg mqtt.Message) {
				tenantMsg := TenantMessage{
					Message:     msg,
					SavePayload: tenant.SavePayloads,
					Tenant:      tenant.Name,
				}
				MessageHandler(store, tenantMsg)
			}

			if token := client.Subscribe(topic, 0, wrappedHandler); token.Wait() && token.Error() != nil {
				return nil, token.Error()
			}
			log.Info("Subscribed to topic", "id", topic, "tenant", tenant.Name)
		}
	}

	return client, nil
}

func newTLSConfig(cfg config.MQTTConfiguration) *tls.Config {
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(cfg.TlsCertFile)
	if err != nil {
		log.Fatal("Error reading CA file",
			"file", cfg.TlsCertFile,
			"error", err)
	}
	certPool.AppendCertsFromPEM(ca)
	return &tls.Config{
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS12,
	}
}
