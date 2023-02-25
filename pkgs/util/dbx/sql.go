package dbx

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

func UnwrapMySQLError(err error) (res *mysql.MySQLError) {
	errors.As(err, &res)
	return
}

// https://dev.mysql.com/doc/mysql-errors/5.7/en/server-error-reference.html
const (

	// 1050-1099
	// https://www.fromdual.com/mysql-error-codes-and-messages-1050-1099

	// MySQLErrCodeDuplicateKey
	//
	// Duplicate key name '%s'
	MySQLErrCodeDuplicateKey = 1061

	// MySQLErrCodeDuplicateEntity
	//
	// Duplicate entry '%s' for key %d
	MySQLErrCodeDuplicateEntity = 1062
)
