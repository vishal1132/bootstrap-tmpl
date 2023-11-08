{% if redis_enabled %}
package status

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisMonitor struct {
	name string
	conn *redis.Client
}

type option func(*RedisMonitor)

func NewRedisMonitor(name string, conn *redis.Client, opts ...option) *RedisMonitor {
	m := &RedisMonitor{name: name, conn: conn}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *RedisMonitor) GetName() string {
	return m.name
}

func (m *RedisMonitor) Ping(ctx context.Context) error {
	return logDependencyError(ctx, "Redis", m.name, m.conn.Ping(ctx).Err())
}

{% endif %}