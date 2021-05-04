package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/controllers"
	"github.com/riquefsouza/hefesto-fiber-gorm/database"
)

var URL = "/api/v1"

func main() {
	app := fiber.New()

	app.Static("/", "public")

	database.InitDatabase()
	defer database.DBConn.Close()

	controllers.Routes_AdmParameterCategory(app, URL)
	controllers.Routes_AdmParameter(app, URL)

	log.Fatal(app.Listen(":3000"))
}
