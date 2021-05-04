package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riquefsouza/hefesto-fiber-gorm/services"
)

func Routes_AdmParameterCategory(app *fiber.App, URL string) {
	app.Get(URL+"/admParameterCategory", services.AdmParameterCategory_FindAll)
	app.Get(URL+"/admParameterCategory/:id", services.AdmParameterCategory_FindById)
	app.Post(URL+"/admParameterCategory", services.AdmParameterCategory_Insert)
	app.Put(URL+"/admParameterCategory/:id", services.AdmParameterCategory_Update)
	app.Delete(URL+"/admParameterCategory/:id", services.AdmParameterCategory_Delete)
}
