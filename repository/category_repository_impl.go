package repository

import (
	"context"
	"database/sql"
	"errors"
	"restapi-golang/helper"
	"restapi-golang/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepositoryImpl{}
}

func (CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	helper.PanicError(err)
	category.Id = int(id)
	return category
}

func (repository CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicError(err)
	//TODO implement me
	return category
}

func (repository CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicError(err)
	//TODO implement me
}

func (repository CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	//TODO implement me
	SQL := "select id, name from category  where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicError(err)
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	//TODO implement me
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&categories, category)
		helper.PanicError(err)
		categories = append(categories, category)
	}
	return categories
}
