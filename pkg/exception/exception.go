package exception

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	logrus.Error(statusDetails.Details())
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

func UnauthenticatedError() error {
	return status.Errorf(codes.Unauthenticated, "unauthorized")
}

func InternalError() error {
	return status.Errorf(codes.Internal, "internal")
}

func DatabaseQueryError() error {
	return status.Errorf(codes.Canceled, "database query error")
}

func UnknownError() error {
	return status.Errorf(codes.Unknown, "unknown")
}
