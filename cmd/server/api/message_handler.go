package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"gorm.io/gorm"
)

func MessageHandler(conn *gorm.DB, c *fiber.Ctx) error {
	// Parsing dates from the request parameters
	startDateStr := c.Query("start_datetime")
	endDateStr := c.Query("end_datetime")

	// Error handling for date parsing
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_date format"})
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_date format"})
	}

	// Using the functions to get the data
	totalMessages, err := db.GetTotalMessages(conn, startDate, endDate)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dailyMessagesPerTenant, err := db.GetDailyMessagesPerTenant(conn, startDate, endDate)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"total_messages": totalMessages,
		"daily_data":     dailyMessagesPerTenant,
	})
}
