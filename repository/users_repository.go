package repository

import (
	"context"
	"database/sql"
	"restapi-golang/model/domain"
)

type UsersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	// Update(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	// FindById(ctx context.Context, tx *sql.Tx, usersId int) domain.Users
	// FindAll(ctx context.Context, tx *sql.Tx) []domain.Users
}
