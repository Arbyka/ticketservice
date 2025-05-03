package controller

import (
	"net/http"
	"strconv"
	"ticketing/service"

	"github.com/gin-gonic/gin"
)

type ReportsController struct {
	service service.ReportsService
}

func NewReportsController(service service.ReportsService) *ReportsController {
	return &ReportsController{service}
}

func (c *ReportsController) GetSummaryReport(ctx *gin.Context) {
	totalTickets, totalRevenue, err := c.service.GetSummary()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_tickets_sold": totalTickets,
		"total_revenue":      totalRevenue,
	})
}

func (c *ReportsController) GetEventReport(ctx *gin.Context) {
	idParam := ctx.Param("id")
	eventID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	ticketsSold, totalRevenue, err := c.service.GetEventReport(uint(eventID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"event_id":           eventID,
		"tickets_sold":       ticketsSold,
		"event_revenue":      totalRevenue,
	})
}
