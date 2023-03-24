package pkg

import db "github.com/nguyenkhoa0721/go-project-layout/pkg/db/postgres/sqlc"

type GetChainResponse struct {
	db.Chain
}

type GetManyChainsResponse struct {
	Rows []db.Chain
}
