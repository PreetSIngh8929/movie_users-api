package mysql_utils

import (
	"errors"
	"strings"

	"github.com/PreetSIngh8929/movie_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("Invalid Data")
	}
	return rest_errors.NewInternalServerError("error parsing request", errors.New("database error"))
}
