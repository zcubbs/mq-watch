package mock

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
	"github.com/zcubbs/x/must"
)

var (
	broker string
	topics []string
)

var Cmd = &cobra.Command{
	Use:   "mock",
	Short: "mock",
	Long:  "example: mq-watch mock -b tcp://localhost:1883 -t tenant1/topic1 -t tenant2/topic1",
	Run: func(cmd *cobra.Command, args []string) {
		must.Succeed(mock())
	},
}

func mock() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("mq-watch-cli-mock")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	defer client.Disconnect(250)

	publishMockData(client, topics)

	return nil
}

func init() {
	Cmd.Flags().StringVarP(&broker, "broker", "b", "tcp://localhost:1883", "kafka broker")
	Cmd.Flags().StringArrayVarP(&topics, "topics", "t", nil, "topics to mock")

	_ = Cmd.MarkFlagRequired("topics")
}
