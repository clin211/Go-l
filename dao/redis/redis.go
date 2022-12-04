package redis

import (
	"fmt"
	"go-web/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.Conf.Redis.Host,
			settings.Conf.Redis.Port,
		),
		Password: settings.Conf.Redis.Password,
		DB:       settings.Conf.Redis.Db,
		PoolSize: settings.Conf.Redis.PoolSize,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
