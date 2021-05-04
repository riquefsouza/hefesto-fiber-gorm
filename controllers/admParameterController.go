package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/services"
)

func Routes_AdmParameter(app *fiber.App, URL string) {
	app.Get(URL+"/admParameter", services.AdmParameter_FindAll)
	app.Get(URL+"/admParameter/:id", services.AdmParameter_FindById)
	app.Post(URL+"/admParameter", services.AdmParameter_Insert)
	app.Put(URL+"/admParameter/:id", services.AdmParameter_Update)
	app.Delete(URL+"/admParameter/:id", services.AdmParameter_Delete)
}
