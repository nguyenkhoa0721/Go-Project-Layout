package pkg

import (
	"github.com/nguyenkhoa0721/go-project-layout/pkg/exception"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/validator"
	pb "github.com/nguyenkhoa0721/grpc/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func ValidateGetChain(req *pb.GetChainRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateNumeric(req.Id); err != nil {
		violations = append(violations, exception.FieldViolation("id", err))
	}

	return violations
}
