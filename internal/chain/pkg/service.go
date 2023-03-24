package pkg

import (
	"context"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/common"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/exception"
)

type ChainService struct {
	common *common.Common
}

func NewChainService(c *common.Common) *ChainService {
	return &ChainService{
		common: c,
	}
}

func (service *ChainService) GetChain(id string) (*GetChainResponse, error) {
	chain, err := service.common.Store.GetChain(context.Background(), id)
	if err != nil {
		return nil, exception.DbException(err)
	}

	return &GetChainResponse{
		Chain: chain,
	}, nil
}

func (service *ChainService) GetManyChains() (*GetManyChainsResponse, error) {
	chains, err := service.common.Store.GetManyChain(context.Background())
	if err != nil {
		return nil, exception.DbException(err)
	}

	return &GetManyChainsResponse{
		chains,
	}, nil
}
