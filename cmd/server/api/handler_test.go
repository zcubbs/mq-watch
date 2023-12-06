package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zcubbs/mq-watch/cmd/server/config"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"net/http/httptest"
	"testing"
)

func NewMockServerConfiguration() config.ServerConfiguration {
	return config.ServerConfiguration{
		Port:        8080,
		TlsEnabled:  false,
		TlsCertFile: "",
		TlsKeyFile:  "",
	}
}

func TestMessageHandler(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Get("/api/messages", func(c *fiber.Ctx) error {
		return messageHandler(store, c)
	})

	req := httptest.NewRequest("GET", "/api/messages?start_datetime=2022-01-01T00:00:00Z&end_datetime=2022-01-31T23:59:59Z", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestTotalMessagesPerDayHandler(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Get("/api/total-messages-per-day", func(c *fiber.Ctx) error {
		return totalMessagesPerDayHandler(store, c)
	})

	req := httptest.NewRequest("GET", "/api/total-messages-per-day?start_date=2022-01-01&end_date=2022-01-31", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestHealthHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/health", healthHandler)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestMountApiRoutes(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Get("/api/messages", func(c *fiber.Ctx) error {
		return messageHandler(store, c)
	})
	app.Get("/api/total-messages-per-day", func(c *fiber.Ctx) error {
		return totalMessagesPerDayHandler(store, c)
	})
	app.Get("/health", healthHandler)

	req := httptest.NewRequest("GET", "/api/messages?start_datetime=2022-01-01T00:00:00Z&end_datetime=2022-01-31T23:59:59Z", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	req = httptest.NewRequest("GET", "/api/total-messages-per-day?start_date=2022-01-01&end_date=2022-01-31", nil)
	resp, _ = app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	req = httptest.NewRequest("GET", "/health", nil)
	resp, _ = app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestMountStaticRoutes(t *testing.T) {
	app := fiber.New()
	app.Get("/health", healthHandler)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestNewServer(t *testing.T) {
	store := db.NewMockStore()
	cfg := NewMockServerConfiguration()

	server := NewServer(cfg, store)

	assert.Equal(t, store, server.store)
	assert.Equal(t, cfg, server.cfg)
}

func TestGetMessageStatsHandler(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Get("/api/message-stats", func(c *fiber.Ctx) error {
		return getMessageStatsHandler(store, c)
	})

	req := httptest.NewRequest("GET", "/api/message-stats", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetTopTenantsHandler(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Get("/api/top-tenants", func(c *fiber.Ctx) error {
		return getTopTenantsHandler(store, c)
	})

	req := httptest.NewRequest("GET", "/api/top-tenants", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestSaveMessagesHandler(t *testing.T) {
	store := db.NewMockStore()

	app := fiber.New()
	app.Post("/api/save-messages", func(c *fiber.Ctx) error {
		return saveMessagesHandler(store, c)
	})

	req := httptest.NewRequest("POST", "/api/save-messages", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
