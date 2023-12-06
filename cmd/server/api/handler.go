package api

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/mq-watch/cmd/server/db"
)

// HealthHandler returns a simple "OK" string to indicate the server is running.
func healthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

// messageHandler retrieves the total messages count for a date range and optional tenant filter.
func messageHandler(store db.Store, c *fiber.Ctx) error {
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
	totalMessages, err := store.GetTotalMessages(startDate, endDate)
	if err != nil {
		log.Error("Error getting total messages", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dailyMessagesPerTenant, err := store.GetDailyMessagesPerTenant(startDate, endDate)
	if err != nil {
		log.Error("Error getting daily messages per tenant", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"total_messages": totalMessages,
		"daily_data":     dailyMessagesPerTenant,
	})
}

// totalMessagesPerDayHandler retrieves the total messages count per day for a date range.
func totalMessagesPerDayHandler(store db.Store, c *fiber.Ctx) error {
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
	messagesTotalPerDay, err := store.GetMessagesTotalPerDay(startDate, endDate)
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
		parsedDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			log.Error("Error parsing date", "date", date, "error", err)
			continue
		}
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

// getTopTenantsHandler retrieves the top tenants based on message count.
func getTopTenantsHandler(store db.Store, c *fiber.Ctx) error {
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
	topTenants, err := store.GetTopTenants(startDate, endDate)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Return the top tenants in JSON format
	return c.Status(http.StatusOK).JSON(topTenants)
}

// getMessageStatsHandler retrieves the total messages count for a date range and optional tenant filter.
func getMessageStatsHandler(store db.Store, c *fiber.Ctx) error {
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
		totalMessages, err = store.GetMessagesPerTenant(tenant, startDate, endDate)
	} else {
		// Fetch total messages across all tenants within the date range
		totalMessages, err = store.GetTotalMessages(startDate, endDate)
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"total_messages": totalMessages})
}

type saveMessageRequest struct {
	Tenant    string `json:"tenant"`
	Topic     string `json:"topic"`
	Payload   string `json:"payload"`
	CreatedAt string `json:"created_at"`
}

type saveMessagesRequest struct {
	Messages []saveMessageRequest `json:"messages"`
}

// saveMessagesHandler saves messages to the database.
func saveMessagesHandler(store db.Store, c *fiber.Ctx) error {
	var req saveMessagesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	for _, msg := range req.Messages {
		if err := store.SaveMessage(msg.Tenant, msg.Topic, msg.Payload, msg.CreatedAt); err != nil {
			log.Error("Error saving message", "tenant", msg.Tenant, "topic", msg.Topic, "error", err)
			return c.Status(http.StatusInternalServerError).JSON(
				fiber.Map{"error": fmt.Sprintf("Error saving message: %v", err)},
			)
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Messages saved successfully"})
}
