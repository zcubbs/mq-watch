package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/zcubbs/x/must"
)

type Message struct {
	Tenant    string `json:"tenant"`
	Topic     string `json:"topic"`
	Payload   string `json:"payload"`
	CreatedAt string `json:"created_at"`
}

type SaveMessagesRequest struct {
	Messages []Message `json:"messages"`
}

var (
	apiURL string
	tenant string
	topic  string
	count  int
	date   string
)

var Cmd = &cobra.Command{
	Use:   "mock",
	Short: "mock",
	Long:  "example: mqw mock -api http://localhost:8000 -t tenant1/topic1 -c 100 -d \"2023-03-15T15:04:05Z\"",
	Run: func(cmd *cobra.Command, args []string) {
		must.Succeed(mock())
	},
}

func mock() error {
	messages := generateMockMessages(tenant, topic, count, date)
	return sendMessagesToAPI(apiURL, messages)
}

func generateMockMessages(tenant, topic string, count int, date string) []Message {
	var messages []Message
	for i := 0; i < count; i++ {
		messages = append(messages, Message{
			Payload:   fmt.Sprintf("Data_%d", i),
			Tenant:    tenant,
			Topic:     topic,
			CreatedAt: date,
		})
	}
	return messages
}

func sendMessagesToAPI(url string, messages []Message) error {
	requestBody, err := json.Marshal(SaveMessagesRequest{Messages: messages})
	if err != nil {
		return err
	}

	resp, err := http.Post(url+"/api/save-messages", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	fmt.Println("API request succeeded")

	return nil
}

func init() {
	Cmd.Flags().StringVarP(&apiURL, "api", "a", "http://localhost:8000", "API URL")
	Cmd.Flags().StringVarP(&tenant, "tenant", "T", "", "tenant name")
	Cmd.Flags().StringVarP(&topic, "topic", "t", "", "topic to publish to")
	Cmd.Flags().IntVarP(&count, "count", "c", 100, "number of messages to publish")
	Cmd.Flags().StringVarP(&date, "date", "d", time.Now().Format(time.RFC3339), "date to publish messages")

	_ = Cmd.MarkFlagRequired("tenant")
	_ = Cmd.MarkFlagRequired("topic")
}
