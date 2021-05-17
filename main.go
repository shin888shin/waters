package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shin888shin/waters/database"
	"github.com/shin888shin/waters/lake"
)

const (
	DBUser     = "postgres"
	DBPassword = "secret"
	DBName     = "theplanet"
	DBHost     = "db"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lake", lake.GetLakes)
	app.Get("/api/v1/lake/:id", lake.GetLake)
	app.Post("/api/v1/lake", lake.NewLake)
	app.Delete("/api/v1/lake/:id", lake.DeleteLake)
}

func initDatabase() {
	var err error
	strDb := GetPostgresConnectionString()
	// database.DBConn, err = gorm.Open("sqlite3", "waters.db")
	database.DBConn, err = gorm.Open("postgres", strDb)
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
