package lake

import (
	"github.com/gofiber/fiber"
)

func GetLakes(c *fiber.Ctx) {
	c.Send("All lakes")
}

func GetLake(c *fiber.Ctx) {
	c.Send("A single lake")
}
func NewLake(c *fiber.Ctx) {
	c.Send("Add a lake")
}
func DeleteLake(c *fiber.Ctx) {
	c.Send("Delete a lake")
}
