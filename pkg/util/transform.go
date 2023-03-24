package util

import (
	"database/sql"
	"errors"
	"strconv"
)

func NullStringToString(val sql.NullString) string {
	if val.Valid {
		return val.String
	} else {
		return ""
	}
}

func StrToInt64(val string) (int64, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return i, nil
}

func StrToInt32(val string) (int32, error) {
	i, err := StrToInt64(val)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return int32(i), nil
}
