package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/services"
)

func Routes_AdmParameterCategory(app *fiber.App, URL string) {
	app.Get(URL+"/admParameterCategory", services.FindAll)
	app.Get(URL+"/admParameterCategory/:id", services.FindById)
	app.Post(URL+"/admParameterCategory", services.Insert)
	app.Put(URL+"/admParameterCategory/:id", services.Update)
	app.Delete(URL+"/admParameterCategory/:id", services.Delete)
}
