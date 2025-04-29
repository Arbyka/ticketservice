package main

import (
    "ticketing/config"
    "ticketing/controller"
    "ticketing/repository"
    "ticketing/service"
    "ticketing/route"
)

func main() {
    config.ConnectDatabase()

    db := config.DB
	userRepo := repository.NewUserRepository(db)
    authService := service.NewAuthService(userRepo)
    authController := controller.NewAuthController(authService)

    r := route.SetupRouter(authController)
    r.Run(":8080")
}
