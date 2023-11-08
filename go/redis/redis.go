{% if redis_enabled %}
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Operation = string

const (
	operationGet                    = "get"
	operationMget                   = "mget"
	operationSet                    = "set"
	operationEvict                  = "evict"
	operationZrangeByScoreWithScore = "zrangeByScoreWithScore"
	operationZadd                   = "zadd"
	operationExpire                 = "expire"
	zremOperationKey                = "zrem"
	zremRangeByScoreOperationKey    = "zremRangeByScore"
)

type RedisDaoImpl struct {
	Rc *redis.Client
}

type IRedisDao interface {
	Get(ctx context.Context, key, segment string) (string, error)
	Set(ctx context.Context, key, segment string, value interface{}, expiration time.Duration) error
	Evict(ctx context.Context, key, segment string) error
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZRemRangeByScore(ctx context.Context, key string, initialScore, endScore int64) error
	ZRangeByScoreWithScores(ctx context.Context, key string, start int64, count int64,
		minScore string, maxScore string) ([]redis.Z, error)
	ZAdd(ctx context.Context, key string, members []redis.Z) (int64, error)
	Expire(ctx context.Context, key string, time time.Duration) (bool, error)
	Mget(ctx context.Context, keys []string, segment string) ([]interface{}, error)
}

func NewRedisDaoImpl(redisClient *redis.Client) IRedisDao {
	return &RedisDaoImpl{
		Rc: redisClient,
	}
}

func (r *RedisDaoImpl) Get(ctx context.Context, key, segment string) (string, error) {
	// defer ctx.Metric.StartRedisSegment(segment, operationGet).End()
	return r.Rc.Get(ctx, key).Result()
}

func (r *RedisDaoImpl) Mget(ctx context.Context, keys []string, segment string) ([]interface{}, error) {
	// defer ctx.Metric.StartRedisSegment(segment, operationMget).End()
	return r.Rc.MGet(ctx, keys...).Result()
}

func (r *RedisDaoImpl) Set(ctx context.Context, key, segment string, value interface{}, expiration time.Duration) error {
	// defer ctx.Metric.StartRedisSegment(segment, operationSet).End()
	if val, ok := value.(string); ok {
		return r.Rc.Set(ctx, key, val, expiration).Err()
	}
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Rc.Set(ctx, key, string(b), expiration).Err()
}

func (r *RedisDaoImpl) Evict(ctx context.Context, key, segment string) error {
	// defer ctx.Metric.StartRedisSegment(segment, operationEvict).End()
	return r.Rc.Del(ctx, key).Err()
}

func (r *RedisDaoImpl) ZRangeByScoreWithScores(ctx context.Context, key string, start int64, count int64,
	minScore string, maxScore string) ([]redis.Z, error) {
	// defer ctx.Metric.StartRedisSegment(operationZrangeByScoreWithScore, operationZrangeByScoreWithScore).End()
	return r.Rc.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    minScore,
		Max:    maxScore,
		Offset: start,
		Count:  count,
	}).Result()
}

func (r *RedisDaoImpl) ZAdd(ctx context.Context, key string, members []redis.Z) (int64, error) {
	// defer ctx.Metric.StartRedisSegment(operationZadd, operationZadd).End()
	return r.Rc.ZAdd(ctx, key, members...).Result()
}

func (r *RedisDaoImpl) Expire(ctx context.Context, key string, time time.Duration) (bool, error) {
	// defer ctx.Metric.StartRedisSegment(operationExpire, operationExpire).End()
	return r.Rc.Expire(ctx, key, time).Result()
}

func (r *RedisDaoImpl) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	// defer ctx.Metric.StartRedisSegment(zremOperationKey, zremOperationKey).End()
	return r.Rc.ZRem(ctx, key, members...).Result()
}

func (r *RedisDaoImpl) ZRemRangeByScore(ctx context.Context, key string, initialScore, endScore int64) error {
	// defer ctx.Metric.StartRedisSegment(zremRangeByScoreOperationKey, zremRangeByScoreOperationKey).End()
	_, err := r.Rc.ZRemRangeByScore(ctx, key, fmt.Sprint(initialScore), fmt.Sprint(endScore)).Result()
	return err
}
{% endif %}