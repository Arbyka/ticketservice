package route

import (
    "github.com/gin-gonic/gin"
    "ticketing/controller"
	"ticketing/middleware"
)

func SetupRouter(authController *controller.AuthController, eventController *controller.EventController) *gin.Engine {
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
        ticket.GET("",) // Melihat daftar tiket
        ticket.POST("",) // Membeli atau memesan tiket.
        ticket.GET("/:id",) // Melihat detail tiket tertentu.
        ticket.PATCH("/:id",) // Memperbarui status tiket menjadi 'cancelled'.
    }

    return r
}
