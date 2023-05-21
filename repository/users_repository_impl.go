package repository

import (
	"context"
	"database/sql"
	"restapi-golang/helper"
	"restapi-golang/model/domain"
)

type UsersRepositoryImpl struct {
}

func NewUsersRepositoryImpl() UsersRepository {
	return UsersRepositoryImpl{}
}

func (UsersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userDomain *domain.Users) domain.Users {
	SQL := "insert into users(namalengkap, username, email, isPerusahaan, password) values(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, userDomain.UserName, userDomain.Email, userDomain.IsPerusahaan, userDomain.Password)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	helper.PanicError(err)
	userDomain.Id = int64(int(id))
	return userDomain
}

func (*UsersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userDomain domain.Users) domain.Users {
	SQL := "update users set namalengkap = ?, username = ?, email = ?, isPerusahaan = ?, password = ? where id = ?"
	result, err := tx.ExecContext(ctx, SQL, userDomain.UserName, userDomain.Email, userDomain.IsPerusahaan, userDomain.Password)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	helper.PanicError(err)
	userDomain.Id = int64(int(id))
	return userDomain
}
func (*UsersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *domain.Users) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicError(err)
}
