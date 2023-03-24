package exception

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DbException(err error) error {
	logrus.Error(err)

	if err.Error() == "sql: no rows in result set" {
		return status.Errorf(codes.NotFound, "Not found")
	} else {
		return DatabaseQueryError()
	}
}
