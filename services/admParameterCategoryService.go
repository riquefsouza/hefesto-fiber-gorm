package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/database"
	"github.com/riquefsouza/hefesto-fiber-gorm/models"
)

func FindAll(c *fiber.Ctx) error {
	db := database.DBConn
	var list []models.AdmParameterCategory
	db.Find(&list)
	return c.JSON(list)
}

func FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var obj models.AdmParameterCategory
	db.Find(&obj, id)
	return c.JSON(obj)
}

func Insert(c *fiber.Ctx) error {
	db := database.DBConn
	//var obj models.AdmParameterCategory
	obj := new(models.AdmParameterCategory)
	if err := c.BodyParser(obj); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&obj)
	return c.JSON(obj)
}

func Update(c *fiber.Ctx) error {
	msg := fmt.Sprintf("update, %s", c.Params("id"))
	return c.SendString(msg)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var obj models.AdmParameterCategory
	db.First(&obj, id)
	if obj.Description == "" {
		return c.Status(500).SendString("No Parameter Category found")
	}
	db.Delete(&obj)
	msg := fmt.Sprintf("deleted with id: %s", id)
	return c.SendString(msg)
}
