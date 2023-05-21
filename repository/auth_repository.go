package repository

import (
	"context"
	"database/sql"
	"restapi-golang/model/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, tx sql.Tx, user domain.Users) domain.Users
	Login(ctx context.Context, tx sql.Tx, user domain.Users) domain.Users
}
