package routes

import (
	"github.com/elasticsearch-tutorial/controllers"
	"github.com/gofiber/fiber/v2"
)

func SearchRoutes(app *fiber.App) {
	route := app.Group("/v1/search")

	route.Get("", controllers.Search)
}
