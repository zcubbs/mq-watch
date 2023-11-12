package main

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
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
	for _, tenant := range tenants {
		// Generate 1 to 3 topics for each tenant, but ensure Tenant1 has exactly 3 topics.
		numTopics := int(getRandInt(3) + 1)
		if tenant == "Tenant1" {
			numTopics = 3
		}

		for i := 0; i < numTopics; i++ {
			topic := fmt.Sprintf("%s/topic%d", tenant, i+1)
			go func(tenant, topic string) {
				for {
					// Generate a random CreatedAt date within the last 7 days
					daysAgo := int(getRandInt(7) + 1)
					createdAt := time.Now().AddDate(0, 0, -daysAgo).Format(time.RFC3339)

					// Create an instance of MessagePayload with mock data
					payload := MessagePayload{
						Device:    fmt.Sprintf("Device_%d", getRandInt(100)),
						Tenant:    tenant,
						CreatedAt: createdAt,
						Data:      fmt.Sprintf("Data_%d", getRandInt(100)),
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

func getRandInt(max int64) int64 {
	randInt, err := crand.Int(crand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Printf("Error generating random int: %v", err)
	}

	return randInt.Int64()
}
