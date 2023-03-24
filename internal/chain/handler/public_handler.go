package handler

import (
	"context"

	"github.com/nguyenkhoa0721/go-project-layout/internal/chain/pkg"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/exception"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/util"
	pb "github.com/nguyenkhoa0721/grpc/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ChainPublicHandler struct {
	pb.UnimplementedChainServer
	service *pkg.ChainService
}

func NewChainPublicHandler() *ChainPublicHandler {
	return &ChainPublicHandler{
		service: pkg.NewChainService(),
	}
}

func (handler *ChainPublicHandler) GetChain(ctx context.Context, request *pb.GetChainRequest) (*pb.GetChainResponse, error) {
	violations := pkg.ValidateGetChain(request)
	if violations != nil {
		return nil, exception.InvalidArgumentError(violations)
	}

	chain, err := handler.service.GetChain(request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetChainResponse{
		Name:   util.NullStringToString(chain.Name),
		Symbol: util.NullStringToString(chain.Symbol),
	}, nil
}

func (handler *ChainPublicHandler) GetManyChains(ctx context.Context, request *emptypb.Empty) (*pb.GetManyChainsResponse, error) {
	chains, err := handler.service.GetManyChains()
	if err != nil {
		return nil, err
	}

	res := make([]*pb.GetChainResponse, len(chains.Rows))
	for i, v := range chains.Rows {
		res[i] = &pb.GetChainResponse{
			Name:   util.NullStringToString(v.Name),
			Symbol: util.NullStringToString(v.Symbol),
		}
	}

	return &pb.GetManyChainsResponse{
		Rows: res,
	}, nil
}
