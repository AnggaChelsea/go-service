package app

import (
	"database/sql"
	"restapi-golang/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:Sayangmamah1.@tcp(localhost:3306)/belajar_golang_restapi")
	helper.PanicError(err)
	db.SetMaxOpenConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third

	//migrate -database "mysql://root:Sayangmamah1.@tcp(localhost:3306)/belajar_golang_restapi" -path db/migrations up
	//migrate -database "mysql://root:Sayangmamah1.@tcp(localhost:3306)/belajar_golang_restapi" -path db/migrations up 2

}
