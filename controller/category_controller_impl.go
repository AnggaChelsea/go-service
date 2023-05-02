package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi-golang/helper"
	"restapi-golang/model/web"
	"restapi-golang/services"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.JsonFormBodyRequest(request, &categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryCreateRequest,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.JsonFormBodyRequest(request, &categoryUpdateRequest)
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.Updated(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)
	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)
	responseFindById := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseFindById,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	responseFindAll := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   responseFindAll,
	}
	helper.WriteToResponseBody(writer, webResponse)

}
