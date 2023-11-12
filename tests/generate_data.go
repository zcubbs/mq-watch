package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MessagePayload struct {
	Device    string `json:"device"`
	Tenant    string `json:"tenant"`
	CreatedAt string `json:"created_at"`
	Data      string `json:"data"`
}

const (
	broker   = "tcp://127.0.0.1:1883"
	clientID = "mock-data-generator"
)

var tenants = []string{"Tenant1", "Tenant2", "Tenant3", "Tenant4", "Tenant5", "Tenant6", "Tenant7", "Tenant8", "Tenant9", "Tenant10"}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	publishMockData(client)
}

func publishMockData(client mqtt.Client) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	for _, tenant := range tenants {
		// Generate 1 to 3 topics for each tenant, but ensure Tenant1 has exactly 3 topics.
		numTopics := rand.Intn(3) + 1
		if tenant == "Tenant1" {
			numTopics = 3
		}

		for i := 0; i < numTopics; i++ {
			topic := fmt.Sprintf("%s/topic%d", tenant, i+1)
			go func(tenant, topic string) {
				for {
					// Generate a random CreatedAt date within the last 7 days
					daysAgo := rand.Intn(7)
					createdAt := time.Now().AddDate(0, 0, -daysAgo).Format(time.RFC3339)

					// Create an instance of MessagePayload with mock data
					payload := MessagePayload{
						Device:    fmt.Sprintf("Device_%d", rand.Intn(100)),
						Tenant:    tenant,
						CreatedAt: createdAt,
						Data:      fmt.Sprintf("Data_%d", rand.Intn(100)),
					}

					// Marshal the payload into a JSON string
					jsonPayload, err := json.Marshal(payload)
					if err != nil {
						fmt.Printf("Error marshaling JSON: %v", err)
						continue
					}

					token := client.Publish(topic, 0, false, jsonPayload)
					token.Wait()

					fmt.Printf("Published message to topic %s: %s\n", topic, jsonPayload)

					// Publish every 1 second
					time.Sleep(1 * time.Second)
				}
			}(tenant, topic) // Pass both tenant and topic to the goroutine
		}
	}
	// Keep the program running
	select {}
}
