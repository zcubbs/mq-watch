package db

import (
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/gorm"
	"time"
)

// GetTotalMessages gets the total number of messages between two dates
func GetTotalMessages(db *gorm.DB, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := db.Model(&models.Message{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error

	return count, err
}

// GetMessagesPerTenant gets the total number of messages for a specific tenant between two dates
func GetMessagesPerTenant(db *gorm.DB, tenant string, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := db.Model(&models.Message{}).
		Where("tenant = ? AND created_at BETWEEN ? AND ?", tenant, startDate, endDate).
		Count(&count).Error

	return count, err
}

// GetDailyMessagesPerTenant gets the sum of the count column per day for each tenant between two dates
func GetDailyMessagesPerTenant(db *gorm.DB, startDate time.Time, endDate time.Time) (map[string]map[string]int64, error) {
	type result struct {
		Date   string
		Tenant string
		Count  int64
	}
	var results []result

	err := db.Model(&models.Message{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as date, tenant, COUNT(*) as count").
		Group("DATE(created_at), tenant").
		Order("DATE(created_at) ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	data := make(map[string]map[string]int64)
	for _, r := range results {
		// Parse the Date string into a time.Time object
		parsedDate, err := time.Parse("2006-01-02T15:04:05Z", r.Date)
		if err != nil {
			return nil, err // Handle the error appropriately
		}

		if _, ok := data[parsedDate.Format("2006-01-02")]; !ok {
			data[parsedDate.Format("2006-01-02")] = make(map[string]int64)
		}
		data[parsedDate.Format("2006-01-02")][r.Tenant] = r.Count
	}
	return data, nil
}

// GetMessagesTotalPerDay gets the sum of the count column per day between two dates
func GetMessagesTotalPerDay(db *gorm.DB, startDate time.Time, endDate time.Time) (map[string]int64, error) {
	type result struct {
		Date  string
		Count int64
	}
	var results []result

	err := db.Model(&models.Message{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Group("DATE(created_at)").
		Order("DATE(created_at) ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	data := make(map[string]int64)
	for _, r := range results {
		data[r.Date] = r.Count
	}

	return data, nil
}

// TopTenant is a struct that will hold the tenant name and message count for the top tenants query.
type TopTenant struct {
	Tenant       string `json:"tenant"`
	MessageCount int64  `json:"messageCount"`
}

// GetTopTenants retrieves the top tenants based on message count.
func GetTopTenants(db *gorm.DB, startDate, endDate time.Time) ([]TopTenant, error) {
	var topTenants []TopTenant
	err := db.Model(&models.Message{}).
		Select("tenant, COUNT(*) as message_count").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Group("tenant").
		Order("COUNT(*) DESC").
		Limit(6).
		Find(&topTenants).Error

	if err != nil {
		return nil, err
	}

	return topTenants, nil
}

func SaveMessage(db *gorm.DB, tenant string, topic string, payload string, createDateIn string) error {
	var createDate time.Time
	if createDateIn == "" {
		createDate = time.Now()
	} else {
		createDate, _ = time.Parse(time.RFC3339, createDateIn)
	}

	return db.Create(&models.Message{
		Tenant:    tenant,
		Topic:     topic,
		Payload:   payload,
		CreatedAt: createDate,
	}).Error
}
