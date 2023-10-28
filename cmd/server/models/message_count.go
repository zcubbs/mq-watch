package models

import (
	"time"
)

type MessageCount struct {
	ID        uint `gorm:"primaryKey"`
	Tenant    string
	Count     int
	CreatedAt time.Time
}
