package exception

import (
	"context"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcException(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	result, err := handler(ctx, req)
	log.GrpcLogger(result, err)

	if st, ok := status.FromError(err); !ok {
		statusCode := st.Code()

		if statusCode == codes.Unknown {
			return nil, UnknownError()
		} else if statusCode == codes.Internal {
			return result, InternalError()
		} else if statusCode == codes.Unauthenticated {
			return result, UnauthenticatedError()
		}
	}

	return result, err
}
