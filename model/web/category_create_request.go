package web

type CategoryCreateRequest struct {
	Name string `validate:"max=50,min=1"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"max=50,min=1"`
}
