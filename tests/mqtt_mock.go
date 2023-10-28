package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker = "tcp://localhost:1883"
	topic  = "my_topic"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("mqtt_publisher")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	rand.Seed(time.Now().UnixNano())

	for {
		payload := generatePayload()
		token := client.Publish(topic, 0, false, payload)
		token.Wait()

		time.Sleep(5 * time.Second) // Adjust the sleep time as needed
	}
}

func generatePayload() string {
	device := fmt.Sprintf("device%d", rand.Intn(100))
	tenant := fmt.Sprintf("tenant%d", rand.Intn(10))
	data := fmt.Sprintf("data%d", rand.Intn(1000))

	return fmt.Sprintf(`{"device": "%s", "tenant": "%s", "data": "%s"}`, device, tenant, data)
}
