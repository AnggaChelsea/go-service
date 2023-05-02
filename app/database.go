package app

import (
	"database/sql"
	"restapi-golang/helper"
	"time"
)

func NewDb() *sql.DB {
	dsn := "root:Sayangmamah1.@tcp(localhost:3306)/example_golang_database?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	helper.PanicError(err)
	db.SetMaxOpenConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
