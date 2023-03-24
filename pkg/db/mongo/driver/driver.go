package driver

import (
	"context"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/db/mongo/repo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client             *mongo.Client
	Db                 *mongo.Database
	RawTransactionRepo repo.RawTransactionRepo
}

func NewMongo(uri string, dbName string) (*Mongo, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	logrus.Info("Mongo connected")

	db := client.Database(dbName)

	return &Mongo{
		client,
		db,
		repo.NewRawTransactionRepo(db),
	}, nil
}

func (m *Mongo) CleanUp() {
	m.Client.Disconnect(context.Background())
	logrus.Info("Mongo: Clean up")
}
