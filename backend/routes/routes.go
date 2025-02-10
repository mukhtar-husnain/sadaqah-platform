package routes

import (
	"github.com/mukhtar-husnain/sadaqah-platform/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/register", controllers.RegisterUser)

	// Volunteer request routes
	api.Post("/requests", controllers.CreateVolunteerRequest)
	api.Get("/requests", controllers.GetAllVolunteerRequests)
}
