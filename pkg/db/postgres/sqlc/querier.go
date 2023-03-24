// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateKafkaDeadLetter(ctx context.Context, arg CreateKafkaDeadLetterParams) (KafkaDeadLetterQueue, error)
	GetChain(ctx context.Context, id string) (Chain, error)
	GetManyChain(ctx context.Context) ([]Chain, error)
}

var _ Querier = (*Queries)(nil)