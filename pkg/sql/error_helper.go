package sql

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicateEntry(err error) bool {
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return true
		}
	}

	return false
}
