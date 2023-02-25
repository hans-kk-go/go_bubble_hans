package redis

import (
	"awesomeProject/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	_, err = rdb.Ping().Result()
	return err

}

func Close() {
	_ = rdb.Close()
}
