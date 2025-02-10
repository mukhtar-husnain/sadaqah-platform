package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/mukhtar-husnain/sadaqah-platform/db"
	"github.com/mukhtar-husnain/sadaqah-platform/routes"
)

func main() {
	fmt.Println("Starting backend...")
	db.ConnectDB()

	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
