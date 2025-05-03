package route

import (
    "github.com/gin-gonic/gin"
    "ticketing/controller"
	"ticketing/middleware"
)

func SetupRouter(
    authController *controller.AuthController,
    eventController *controller.EventController,
    ticketController *controller.TicketController,
    reportController *controller.ReportsController,
    ) *gin.Engine {

    r := gin.Default()

    r.POST("/register", authController.Register)
    r.POST("/login", authController.Login)

    event := r.Group("/events")
    {
        event.GET("", eventController.GetEvents)
        event.POST("", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), eventController.CreateEvent)
        event.PUT("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), eventController.UpdateEvent)
        event.DELETE("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), eventController.DeleteEvent)
    }

    ticket := r.Group("/tickets")
	{
		ticket.GET("", ticketController.GetAll)
		ticket.POST("", ticketController.Create)
		ticket.GET("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), ticketController.GetByID)
		ticket.PATCH("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), ticketController.Cancel)
	}

    report := r.Group("/reports", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware())
	{
		report.GET("/summary", reportController.GetSummaryReport)
		report.GET("/event/:id", reportController.GetEventReport)
	}

    return r
}
