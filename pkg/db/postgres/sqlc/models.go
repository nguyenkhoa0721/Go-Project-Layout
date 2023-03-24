// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"

	"github.com/tabbed/pqtype"
)

type Chain struct {
	ID        string                `json:"id"`
	Name      sql.NullString        `json:"name"`
	Symbol    sql.NullString        `json:"symbol"`
	Rpc       pqtype.NullRawMessage `json:"rpc"`
	UpdatedAt sql.NullTime          `json:"updated_at"`
	CreatedAt sql.NullTime          `json:"created_at"`
}

type KafkaDeadLetterQueue struct {
	ID        string       `json:"id"`
	Topic     string       `json:"topic"`
	Value     string       `json:"value"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}