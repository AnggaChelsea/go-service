package services

import (
	"context"
	"restapi-golang/model/web"
)

type UsersService interface {
	Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse
	Delete(ctx context.Context, id int)
}
