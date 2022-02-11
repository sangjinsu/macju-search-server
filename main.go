package main

import (
	"github.com/elasticsearch-tutorial/middleware"
	"github.com/elasticsearch-tutorial/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(0)

	app := fiber.New()

	middleware.FiberMiddleware(app)
	routes.SearchRoutes(app)

	log.Fatal(app.Listen(":8082"))
}
