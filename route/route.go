package route

import (
    "github.com/gin-gonic/gin"
    "ticketing/controller"
)

func SetupRouter(authController *controller.AuthController) *gin.Engine {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/register", authController.Register)
    r.POST("/login", authController.Login)

    return r
}
