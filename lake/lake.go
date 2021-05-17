package lake

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/shin888shin/waters/database"
)

type Lake struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func GetLakes(c *fiber.Ctx) {
	db := database.DBConn
	var lakes []Lake
	db.Find(&lakes)
	c.JSON(lakes)
}

func GetLake(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lake Lake
	result := db.Find(&lake, id)
	if result.Error != nil {
		c.Status(500).Send("No record found")
		return
	}

	c.JSON(lake)
}

func NewLake(c *fiber.Ctx) {
	db := database.DBConn

	lake := new(Lake)
	if err := c.BodyParser(lake); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&lake)
	c.JSON(lake)
}

func DeleteLake(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lake Lake
	db.First(&lake, id)
	if lake.Name == "" {
		c.Status(500).Send("No record found")
		return
	}
	db.Delete(&lake)
	c.Send("Deleted successfully")
}
