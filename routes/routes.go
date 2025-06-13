package routes

import (
	"encode/modules/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Post("/encode", controllers.Encode)
}
