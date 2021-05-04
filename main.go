package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/controllers"
)

var URL = "/api/v1"

func main() {
	app := fiber.New()

	controllers.Routes_AdmParameterCategory(app, URL)

	log.Fatal(app.Listen(":3000"))
}
