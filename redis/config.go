package redis

type Config struct {
	Addr string `env:"REDIS_ADDR"`
}
