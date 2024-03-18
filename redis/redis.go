package redis

import (
	"context"
	pkgConfig "github.com/fobus1289/ufa_shared/config"
	"github.com/fobus1289/ufa_shared/utils"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisService interface {
	SetWithTTL(key int64, value interface{}, timeOut time.Duration) error
	Set(key int64, value interface{}) error
	Get(key int64) interface{}
}

type redisService struct {
	redisClient *redis.Client
}

func NewRedisService() RedisService {
	config := pkgConfig.Load(&config{})
	redisClient := connect(config)
	return &redisService{
		redisClient: redisClient,
	}
}

func connect(config *config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

func (s *redisService) SetWithTTL(key int64, value interface{}, timeOut time.Duration) error {
	err := s.redisClient.Set(context.Background(), utils.Int64ToString(key), value, timeOut).Err()
	return err
}

func (s *redisService) Set(key int64, value interface{}) error {
	err := s.SetWithTTL(key, value, 0)
	return err
}

func (s *redisService) Get(key int64) interface{} {
	val, err := s.redisClient.Get(context.Background(), utils.Int64ToString(key)).Result()
	if err != nil {
		panic(err)
	}

	return val
}
