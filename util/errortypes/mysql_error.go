package errortypes

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

const (
	MySQLDuplicateKey = 1062
)

func IsMySqlError(err error, errId uint16) bool {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) && mysqlError.Number == errId {
		return true
	}
	return false
}
