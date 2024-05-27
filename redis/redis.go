package redis

import (
	"context"
	"time"

	"github.com/fobus1289/ufa_shared/utils"
	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	SetWithTTL(key int64, value any, timeOut time.Duration) error
	Set(key int64, value any) error
	Get(key int64) (any, error)
}

type redisService struct {
	redisClient *redis.Client
}

func NewRedisService(config *Config) RedisService {
	redisClient := connect(config)
	return &redisService{
		redisClient: redisClient,
	}
}

func connect(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

func (s *redisService) SetWithTTL(key int64, value any, timeOut time.Duration) error {
	status := s.redisClient.Set(context.Background(), utils.Int64ToString(key), value, timeOut)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (s *redisService) Set(key int64, value any) error {
	err := s.SetWithTTL(key, value, 0)
	return err
}

func (s *redisService) Get(key int64) (any, error) {
	val, err := s.redisClient.Get(context.Background(), utils.Int64ToString(key)).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}
