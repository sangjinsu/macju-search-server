package middleware

import (
	"github.com/elasticsearch-tutorial/utils/loaddotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func FiberMiddleware(app *fiber.App) {
	allowOrigins :=
		loaddotenv.LoadDotEnv("LOCAL_SERVER")

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: fiber.MethodGet,
	}))
}
