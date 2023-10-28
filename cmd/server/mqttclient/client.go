package mqttclient

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConnectAndSubscribe(broker string, topic string, messageHandler mqtt.MessageHandler) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("error connecting to MQTT broker: %v", token.Error())
	}

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("error subscribing to topic: %v", token.Error())
	}

	return client, nil
}
