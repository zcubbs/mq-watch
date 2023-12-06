package db

import "time"

type MockStore struct {
	TotalMessagesPerDay    map[string]int64
	DailyMessagesPerTenant map[string]map[string]int64
}

func NewMockStore() *MockStore {
	return &MockStore{
		TotalMessagesPerDay: map[string]int64{
			"2021-01-01": 10,
			"2021-01-02": 20,
			"2021-01-03": 30,
		},
		DailyMessagesPerTenant: map[string]map[string]int64{
			"2021-01-01": {
				"tenant1": 10,
				"tenant2": 20,
				"tenant3": 30,
			},
			"2021-01-02": {
				"tenant1": 10,
				"tenant2": 20,
				"tenant3": 30,
			},
			"2021-01-03": {
				"tenant1": 10,
				"tenant2": 20,
				"tenant3": 30,
			},
		},
	}
}

func (s *MockStore) GetTotalMessages(startDate, endDate time.Time) (int64, error) {
	return s.TotalMessagesPerDay[startDate.Format("2006-01-02")], nil
}

func (s *MockStore) GetDailyMessagesPerTenant(startDate, endDate time.Time) (map[string]map[string]int64, error) {
	return s.DailyMessagesPerTenant, nil
}

func (s *MockStore) GetMessagesTotalPerDay(startDate, endDate time.Time) (map[string]int64, error) {
	return s.TotalMessagesPerDay, nil
}

func (s *MockStore) GetTopTenants(startDate, endDate time.Time) ([]TopTenant, error) {
	return []TopTenant{}, nil
}

func (s *MockStore) GetMessagesPerTenant(tenant string, startDate, endDate time.Time) (int64, error) {
	return 0, nil
}

func (s *MockStore) SaveMessage(tenant, topic, payload, createdAt string) error {
	return nil
}
