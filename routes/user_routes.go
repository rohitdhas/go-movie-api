package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohitdhas/movie-api-go/controllers"
)

func UserRoute(app *fiber.App) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetUsers(c)
	})

}
