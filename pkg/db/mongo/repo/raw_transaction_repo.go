package repo

import (
	"context"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/db/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	RawTransaction = "raw_transactions"
)

type RawTransactionRepo struct {
	Db *mongo.Database
}

func NewRawTransactionRepo(Db *mongo.Database) RawTransactionRepo {
	return RawTransactionRepo{
		Db: Db,
	}
}

func (r *RawTransactionRepo) Insert(ctx context.Context, rawTransaction *model.RawTransaction) error {
	marshalByte, err := bson.Marshal(rawTransaction)
	if err != nil {
		return err
	}

	_, err = r.Db.Collection(RawTransaction).InsertOne(ctx, marshalByte)
	if err != nil {
		return err
	}

	return nil
}
