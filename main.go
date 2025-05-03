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

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

    ticketRepo := repository.NewTicketRepository(db)
    ticketService := service.NewTicketService(ticketRepo, eventRepo, db)
    ticketController := controller.NewTicketController(ticketService)

    reportRepo := repository.NewReportsRepository(db)
	reportService := service.NewReportsService(reportRepo)
	reportController := controller.NewReportsController(reportService)

    r := route.SetupRouter(authController, eventController, ticketController, reportController)
    r.Run(":8080")
}
