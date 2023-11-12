package mqttclient

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/logger"
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

func ConnectAndSubscribe(cfg config.MQTTConfiguration, tenants []config.TenantConfiguration, messageHandler mqtt.MessageHandler) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	for _, tenant := range tenants {
		for _, topic := range tenant.Topics {
			if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
				return nil, token.Error()
			}
			log.Info("Subscribed to topic", "id", topic, "tenant", tenant.Name)
		}
	}

	return client, nil
}
