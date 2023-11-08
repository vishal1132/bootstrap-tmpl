{% if redis_enabled %}

package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"{{ module_name }}/config"
	"{{ module_name }}/observability/log"
	"{{ module_name }}/utils"
)

func InitRedis(ctx context.Context, redisConfig *config.RedisConfig) *redis.Client {
	log.Info(ctx, "Connecting to redis")
	address := redisConfig.Host + ":" + redisConfig.Port
	cfg := redis.Options{
		Addr: address,
		DB:   redisConfig.DbId,
	}

	redisClient := redis.NewClient(&cfg)
	_ = utils.Must(redisClient.Ping(ctx).Result())

	log.Info(ctx, "Connected to redis")
	return redisClient
}

{% endif %}