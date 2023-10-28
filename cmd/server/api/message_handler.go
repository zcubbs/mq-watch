package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zcubbs/mq-watch/cmd/server/db"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func MessageHandler(conn *gorm.DB, c *gin.Context) {
	// Parsing dates from the request parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// You might want to add error checking here for the date parsing
	startDate, _ := time.Parse(time.RFC3339, startDateStr)
	endDate, _ := time.Parse(time.RFC3339, endDateStr)

	// Using the functions to get the data
	totalMessages, err := db.GetTotalMessages(conn, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assuming tenant is a parameter you want to optionally filter by
	tenant := c.Query("tenant")
	if tenant != "" {
		messagesPerTenant, err := db.GetMessagesPerTenant(conn, tenant, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"total_messages":      totalMessages,
			"messages_per_tenant": messagesPerTenant,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"total_messages": totalMessages,
		})
	}
}
