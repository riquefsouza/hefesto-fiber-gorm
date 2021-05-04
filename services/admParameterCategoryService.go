package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FindAll(c *fiber.Ctx) error {
	return c.SendString("All")
}

func FindById(c *fiber.Ctx) error {
	msg := fmt.Sprintf("findById, %s", c.Params("id"))
	return c.SendString(msg)
}

func Insert(c *fiber.Ctx) error {
	return c.SendString("post new")
}

func Update(c *fiber.Ctx) error {
	msg := fmt.Sprintf("update, %s", c.Params("id"))
	return c.SendString(msg)
}

func Delete(c *fiber.Ctx) error {
	msg := fmt.Sprintf("delete, %s", c.Params("id"))
	return c.SendString(msg)
}
