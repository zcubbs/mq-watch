package mqttclient

import (
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"time"
)

func MessageHandler(store db.Store, msg TenantMessage) {
	log.Debug("Received message",
		"tenant", msg.Tenant,
		"topic", msg.Message.Topic(),
	)

	var payload string
	if msg.SavePayload {
		payload = string(msg.Message.Payload())
	} else {
		payload = ""
	}

	err := store.SaveMessage(
		msg.Tenant,
		msg.Message.Topic(),
		payload,
		time.Now().Format(time.RFC3339),
	)

	if err != nil {
		log.Error("Error saving message",
			"tenant", msg.Tenant,
			"topic", msg.Message.Topic(),
			"error", err)
	}
}
