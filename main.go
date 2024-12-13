package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohitdhas/movie-api-go/routes"
)

func main() {
	app := fiber.New()

	// Setup Routes
	routes.UserRoute(app)

	// Start server
	app.Listen(":8080")
}
