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
}
