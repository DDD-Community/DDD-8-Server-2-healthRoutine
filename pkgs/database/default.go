package database

import (
	"database/sql"
	"healthRoutine/cmd"
)

func Conn() *sql.DB {
	dbSource := cmd.LoadConfig().DBConn
	db := GetDB(dbSource)
	return db
}
