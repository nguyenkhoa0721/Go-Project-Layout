package common

import (
	"database/sql"
	"fmt"
	"github.com/nguyenkhoa0721/go-project-layout/config"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/db/mongo/driver"
	sqlc "github.com/nguyenkhoa0721/go-project-layout/pkg/db/postgres/sqlc"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/redis"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/uuid"
	"github.com/sirupsen/logrus"
)

var common *Common

type Common struct {
	Config *config.Config
	Store  *sqlc.Queries
	Mongo  *driver.Mongo
	Redis  *redis.Client
	Uuid   *uuid.Uuid
}

func NewCommon() *Common {
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Error(err)
	}

	conn, err := sql.Open(config.Database.Driver, fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.Driver, config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Database,
	))
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	db := sqlc.New(conn)
	logrus.Info("Db connected")

	mongo, err := driver.NewMongo(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.Mongo.Username, config.Mongo.Password, config.Mongo.Host, config.Mongo.Port, config.Mongo.Database), config.Mongo.Database)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	redis, err := redis.NewRedisClient(
		fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port), config.Redis.Password,
	)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	uuid, err := uuid.NewUuid(redis, config.Uuid)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	common = &Common{
		config,
		db,
		mongo,
		redis,
		uuid,
	}
	return common
}

func GetCommon() *Common {
	return common
}
func (c *Common) CleanUp() {
	c.Redis.CleanUp()
	c.Mongo.CleanUp()
	logrus.Info("Common: Clean up")
}
