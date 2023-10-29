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

	// Optional datetime field (can be enabled/disabled based on testing needs)
	var datetimeField string
	// Random date within last 30 days
	datetime := time.Now().Add(time.Duration(-1*rand.Intn(30*24)) * time.Hour)
	datetimeField = fmt.Sprintf(`, "datetime": "%s"`, datetime.Format(time.RFC3339))

	return fmt.Sprintf(`{"device": "%s", "tenant": "%s", "created_at": "%s", "data": "%s"%s}`, device, tenant, datetime.Format(time.RFC3339), data, datetimeField)
}
