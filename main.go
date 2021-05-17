package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shin888shin/waters/database"
	"github.com/shin888shin/waters/lake"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lake", lake.GetLakes)
	app.Get("/api/v1/lake/:id", lake.GetLake)
	app.Post("/api/v1/lake", lake.NewLake)
	app.Delete("/api/v1/lake/:id", lake.DeleteLake)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "waters.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lake.Lake{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(8000)
}
