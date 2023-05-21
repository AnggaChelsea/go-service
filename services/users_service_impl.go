package services

import (
	"context"
	"database/sql"
	"restapi-golang/helper"
	"restapi-golang/model/domain"
	"restapi-golang/model/web"
	"restapi-golang/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UsersRepository repository.UsersRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewUserService(userRepository repository.UsersRepository, DB *sql.DB, validate *validator.Validate) UsersService {
	return &UserServiceImpl{
		UsersRepository: userRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	users := domain.Users{
		NamaLengkap:  request.NamaLengkap,
		UserName:     request.UserName,
		Email:        request.Email,
		IsPerusahaan: request.IsPerusahaan,
		Password:     request.Password,
	}
	users = service.UsersRepository.Save(ctx, tx, users)
	return helper.ToUsersResponse(users)
}

func (service *UserServiceImpl) Delete(ctx context.Context, usersId int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	service.UsersRepository.Delete(ctx, tx, usersId)
}
