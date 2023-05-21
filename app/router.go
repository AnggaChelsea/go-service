package app

import (
	"restapi-golang/controller"
	"restapi-golang/excepction"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.POST("/api/categories", categoryController.Create)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	//users router

	// router.POST("/api/create-users")

	router.PanicHandler = excepction.ErrorHandler
	return router
}
