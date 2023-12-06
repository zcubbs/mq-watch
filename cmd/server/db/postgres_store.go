package db

import (
	"github.com/zcubbs/mq-watch/cmd/server/models"
	"gorm.io/gorm"
	"time"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(db *gorm.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) GetTotalMessages(startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := s.db.Model(&models.Message{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Count(&count).Error

	return count, err
}

func (s *PostgresStore) GetDailyMessagesPerTenant(startDate time.Time, endDate time.Time) (map[string]map[string]int64, error) {
	type result struct {
		Date   string
		Tenant string
		Count  int64
	}
	var results []result

	err := s.db.Model(&models.Message{}).
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

func (s *PostgresStore) GetMessagesTotalPerDay(startDate time.Time, endDate time.Time) (map[string]int64, error) {
	type result struct {
		Date  string
		Count int64
	}
	var results []result

	err := s.db.Model(&models.Message{}).
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

func (s *PostgresStore) GetTopTenants(startDate time.Time, endDate time.Time) ([]TopTenant, error) {
	type result struct {
		Tenant       string
		MessageCount int64
	}
	var results []result

	err := s.db.Model(&models.Message{}).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("tenant, COUNT(*) as message_count").
		Group("tenant").
		Order("message_count DESC").
		Limit(10).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	var topTenants []TopTenant
	for _, r := range results {
		topTenants = append(topTenants, TopTenant{
			Tenant:       r.Tenant,
			MessageCount: r.MessageCount,
		})
	}

	return topTenants, nil
}

func (s *PostgresStore) GetMessagesPerTenant(tenant string, startDate time.Time, endDate time.Time) (int64, error) {
	var count int64
	err := s.db.Model(&models.Message{}).
		Where("tenant = ? AND created_at BETWEEN ? AND ?", tenant, startDate, endDate).
		Count(&count).Error

	return count, err
}

func (s *PostgresStore) SaveMessage(tenant, topic, payload, createdAt string) error {
	var createDate time.Time
	if createdAt == "" {
		createDate = time.Now()
	} else {
		createDate, _ = time.Parse(time.RFC3339, createdAt)
	}

	return s.db.Create(&models.Message{
		Tenant:    tenant,
		Topic:     topic,
		Payload:   payload,
		CreatedAt: createDate,
	}).Error
}
