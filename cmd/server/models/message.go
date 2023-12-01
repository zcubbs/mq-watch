package models

import (
	"time"
)

type Message struct {
	ID        uint `gorm:"primaryKey"`
	Tenant    string
	Topic     string
	Payload   string
	CreatedAt time.Time
}
