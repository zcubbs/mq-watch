package db

import (
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/gorm"
	"time"
)

// GetTotalMessages gets the sum of the count column between two dates
func GetTotalMessages(db *gorm.DB, startDate time.Time, endDate time.Time) (int64, error) {
	var total int64
	err := db.Model(&models.MessageCount{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("COALESCE(SUM(count), 0) as total").
		Scan(&total).Error

	return total, err
}

// GetMessagesPerTenant gets the sum of the count column for a specific tenant between two dates
func GetMessagesPerTenant(db *gorm.DB, tenant string, startDate time.Time, endDate time.Time) (int64, error) {
	var total int64
	err := db.Model(&models.MessageCount{}).
		Where("tenant = ? AND created_at BETWEEN ? AND ?", tenant, startDate, endDate).
		Select("COALESCE(SUM(count), 0) as total").
		Scan(&total).Error

	return total, err
}

// GetDailyMessagesPerTenant gets the sum of the count column per day for each tenant between two dates
func GetDailyMessagesPerTenant(db *gorm.DB, startDate time.Time, endDate time.Time) (map[string]map[string]int64, error) {
	type result struct {
		Date   string // Change this to string
		Tenant string
		Count  int64
	}
	var results []result

	err := db.Model(&models.MessageCount{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as date, tenant, SUM(count) as count").
		Group("DATE(created_at), tenant").
		Order("DATE(created_at) ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	data := make(map[string]map[string]int64)
	for _, r := range results {
		// Parse the Date string into a time.Time object
		parsedDate, err := time.Parse("2006-01-02", r.Date)
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

	err := db.Model(&models.MessageCount{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as date, SUM(count) as count").
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

func SaveMessage(db *gorm.DB, tenant string, count int, createDateIn string) {
	var createDate time.Time
	if createDateIn == "" {
		createDate = time.Now()
	} else {
		createDate, _ = time.Parse(time.RFC3339, createDateIn)
	}

	db.Create(&models.MessageCount{Tenant: tenant, Count: count, CreatedAt: createDate})
}
