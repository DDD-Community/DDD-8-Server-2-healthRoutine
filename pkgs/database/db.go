package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetDB(dbSource string) *sql.DB {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	poolCount := 15
	db.SetMaxIdleConns(poolCount)
	db.SetMaxOpenConns(poolCount)
	db.SetConnMaxLifetime(time.Minute * 3)
	return db
}
