package db

import (
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/gorm"
	"time"
)

// GetTotalMessages get the total message count between two dates
func GetTotalMessages(db *gorm.DB, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := db.Model(&models.MessageCount{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error

	return count, err
}

// GetMessagesPerTenant get the total message count per tenant between two dates
func GetMessagesPerTenant(db *gorm.DB, tenant string, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := db.Model(&models.MessageCount{}).
		Where("tenant = ? AND created_at BETWEEN ? AND ?", tenant, startDate, endDate).
		Count(&count).Error

	return count, err
}
