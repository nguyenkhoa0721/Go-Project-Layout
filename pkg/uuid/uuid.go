package uuid

import (
	"github.com/nguyenkhoa0721/go-project-layout/config"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/redis"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Uuid struct {
	min         uint64
	max         uint64
	nodeCounter uint64
	shardingId  uint64
}

func NewUuid(redisClient *redis.Client, uuidConfig config.Uuid) (*Uuid, error) {
	redisCounter := GetCounter(redisClient, uuidConfig)

	baseCounter := redisCounter % uuidConfig.TotalReplicateNodes
	maxCountInNode := 2048 / uuidConfig.TotalReplicateNodes

	min := baseCounter * maxCountInNode
	max := min + maxCountInNode - 1

	logrus.Infof("Uuid generator MIN: %d, MAX: %d, BASE: %d", min, max, baseCounter)

	return &Uuid{
		min,
		max,
		min,
		uuidConfig.ShardingId,
	}, nil
}

func GetCounter(redisClient *redis.Client, uuidConfig config.Uuid) uint64 {
	var counter uint64 = 0

	strCounter, err := redisClient.GetValue(COUNTER_KEY)
	if err != nil {
		counter = 0
		redisClient.SetValue(COUNTER_KEY, "0")
	}

	counter, err = strconv.ParseUint(strCounter, 10, 32)
	if err != nil {
		counter = 0
		redisClient.SetValue(COUNTER_KEY, "0")
	}

	if counter > COUNTER_MAX {
		redisClient.SetValue(COUNTER_KEY, "0")
	}

	redisClient.IncValue(COUNTER_KEY)

	return counter
}

func (u *Uuid) incNodeCounter() uint64 {
	if u.nodeCounter >= u.max {
		u.nodeCounter = u.min
	}

	u.nodeCounter = u.nodeCounter + 1
	return u.nodeCounter
}

func (u *Uuid) GenerateUuid(uuidType uint64) uint64 {
	counterInt := u.incNodeCounter()
	now := uint64(time.Now().UnixMilli())

	uuid := now << 23
	uuid |= u.shardingId << 17
	uuid |= uuidType << 11
	uuid |= counterInt

	logrus.Infof("New uuid generated: %d. Time: %d, Sharding: %d, Type: %d, Counter: %d", uuid, now, u.shardingId, uuidType, counterInt)
	return uuid
}
