package web

type CategoryCreateRequest struct {
	Name string `validate: "required"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate: "required"`
	Name string `validate: "required"`
}
