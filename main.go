package main

import (
	"github.com/gofiber/fiber"
	"github.com/shin888shin/waters/lake"
)

func hello(c *fiber.Ctx) {
	c.Send("hello")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lake", lake.GetLakes)
	app.Get("/api/v1/lake/:id", lake.GetLake)
	app.Post("/api/v1/lake", lake.NewLake)
	app.Delete("/api/v1/lake/:id", lake.DeleteLake)
}
func main() {
	// fmt.Println("Hello, World!")
	app := fiber.New()
	// app.Get("/", hello)
	setupRoutes(app)
	app.Listen(8000)
}
