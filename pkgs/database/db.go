package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"healthRoutine/pkgs/log"
	"time"
)

const (
	named = "DATABASE_CONN"
)

func GetDB(dbSource string) *sql.DB {
	logger := log.Get()
	defer logger.Sync()

	logger.Named(named).Info("get database connection")
	logger.Named(named).Debug(dbSource)

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
