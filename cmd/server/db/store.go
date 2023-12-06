package db

import "time"

// Store is an interface that includes all the methods you need to mock.
type Store interface {
	GetTotalMessages(startDate, endDate time.Time) (int64, error)
	GetDailyMessagesPerTenant(startDate, endDate time.Time) (map[string]map[string]int64, error)
	GetMessagesTotalPerDay(startDate, endDate time.Time) (map[string]int64, error)
	GetTopTenants(startDate, endDate time.Time) ([]TopTenant, error)
	GetMessagesPerTenant(tenant string, startDate, endDate time.Time) (int64, error)
	SaveMessage(tenant, topic, payload, createdAt string) error
}

// TopTenant is a struct that will hold the tenant name and message count for the top tenants query.
type TopTenant struct {
	Tenant       string `json:"tenant"`
	MessageCount int64  `json:"messageCount"`
}
