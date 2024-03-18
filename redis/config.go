package redis

type config struct {
	Addr string `env:"REDIS_ADDR"`
}
