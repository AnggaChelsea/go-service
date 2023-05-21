package app

import (
	"github.com/julienschmidt/httprouter"
	"restapi-golang/controller"
	"restapi-golang/excepction"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.POST("/api/categories", categoryController.Create)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = excepction.ErrorHandler
	return router
}
