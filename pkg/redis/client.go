package redis

import (
	"fmt"
	"github.com/cespare/xxhash"
	GoRedis "github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	client *GoRedis.Client
}

func NewRedisClient(conn string, password string) (*Client, error) {
	client := GoRedis.NewClient(&GoRedis.Options{
		Addr:     conn,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	logrus.Info("Redis connected")
	return &Client{
		client,
	}, nil
}

func (r *Client) CleanUp() {
	r.client.Close()
	logrus.Info("Redis: Clean up")
}

func (r *Client) CreateKey(prefix string, param *string) string {
	if param == nil {
		return prefix
	}

	hash := xxhash.Sum64([]byte(*param))
	return fmt.Sprintf("%s_%x", prefix, hash)
}

func (r *Client) GetValue(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *Client) SetValue(key string, val string) error {
	err := r.client.Set(key, val, 0).Err()

	return err
}

func (r *Client) SetValueWithExpiry(key string, val string, exp time.Duration) error {
	err := r.client.Set(key, val, exp).Err()

	return err
}

func (r *Client) DeleteKey(key string) error {
	err := r.client.Del(key).Err()

	return err
}

func (r *Client) IncValue(key string) error {
	err := r.client.Incr(key).Err()

	return err
}
