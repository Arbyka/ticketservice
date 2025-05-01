package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ticketing/entity"
	"ticketing/service"
)

type TicketController struct {
	ticketService service.TicketService
}

type OrderItemInput struct {
	EventID  uint `json:"event_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}

type CreateTicketInput struct {
	UserID     uint             `json:"user_id" binding:"required"`
	OrderItems []OrderItemInput `json:"order_items" binding:"required"`
}

func NewTicketController(s service.TicketService) *TicketController {
	return &TicketController{s}
}

// GET /tickets
func (c *TicketController) GetAll(ctx *gin.Context) {
	tickets, err := c.ticketService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

// GET /tickets/:id
func (c *TicketController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	ticket, err := c.ticketService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	ctx.JSON(http.StatusOK, ticket)
}

// POST /tickets
func (c *TicketController) Create(ctx *gin.Context) {
	var input CreateTicketInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Konversi OrderItemInput ke entity.OrderItem
	var orderItems []entity.OrderItem
	for _, item := range input.OrderItems {
		orderItems = append(orderItems, entity.OrderItem{
			EventID:  item.EventID,
			Quantity: item.Quantity,
		})
	}

	ticket := entity.Ticket{
		UserID:     input.UserID,
		OrderItems: orderItems,
	}

	newTicket, err := c.ticketService.Create(ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newTicket)
}

// PATCH /tickets/:id
func (c *TicketController) Cancel(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := c.ticketService.Cancel(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ticket cancelled"})
}
