package api

import (
	"net/http"
	"sort"
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

func TotalMessagesPerDayHandler(conn *gorm.DB, c *fiber.Ctx) error {
	// Parsing dates from the request parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// Error handling for date parsing
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_date format"})
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_date format"})
	}

	// Using the function to get the data
	messagesTotalPerDay, err := db.GetMessagesTotalPerDay(conn, startDate, endDate)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Transforming the map into a slice of the desired structure
	type totalPerDay struct {
		Name  string `json:"name"`
		Total int64  `json:"total"`
	}

	var totals []totalPerDay
	for date, count := range messagesTotalPerDay {
		parsedDate, _ := time.Parse("2006-01-02", date)
		formattedDate := parsedDate.Format("02 Jan") // Format as "DD MMM"
		totals = append(totals, totalPerDay{Name: formattedDate, Total: count})
	}

	// Sorting the slice by date
	sort.Slice(totals, func(i, j int) bool {
		dateI, _ := time.Parse("02 Jan", totals[i].Name)
		dateJ, _ := time.Parse("02 Jan", totals[j].Name)
		return dateI.Before(dateJ)
	})

	// Returning the results in JSON format
	return c.Status(http.StatusOK).JSON(totals)
}

// GetTopTenantsHandler retrieves the top tenants based on message count.
func GetTopTenantsHandler(conn *gorm.DB, c *fiber.Ctx) error {
	// Parsing dates from the request parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// Error handling for date parsing
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_date format"})
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_date format"})
	}

	// Use a database function to get top tenants
	topTenants, err := db.GetTopTenants(conn, startDate, endDate)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Return the top tenants in JSON format
	return c.Status(http.StatusOK).JSON(topTenants)
}

// GetMessageStatsHandler retrieves the total messages count for a date range and optional tenant filter.
func GetMessageStatsHandler(conn *gorm.DB, c *fiber.Ctx) error {
	// Parsing dates from the request parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// Optional tenant filter
	tenant := c.Query("tenant")

	// Error handling for date parsing
	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_date format"})
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_date format"})
	}

	var totalMessages int64
	if tenant != "" {
		// Fetch total messages for a specific tenant within the date range
		totalMessages, err = db.GetMessagesPerTenant(conn, tenant, startDate, endDate)
	} else {
		// Fetch total messages across all tenants within the date range
		totalMessages, err = db.GetTotalMessages(conn, startDate, endDate)
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"total_messages": totalMessages})
}
