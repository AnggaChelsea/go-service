package helper

import (
	"restapi-golang/model/domain"
	"restapi-golang/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func CategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}

func ToUsersResponse(users domain.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:           users.Id,
		NamaLengkap:  users.NamaLengkap,
		Email:        users.Email,
		IsPerusahaan: users.IsPerusahaan,
		Password:     users.Password,
	}
}
