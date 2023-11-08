{% if redis_enabled %}
package config

import "github.com/spf13/viper"

const (
	RedisHost               = "redis.host"
	RedisPort               = "redis.port"
	RedisDbId               = "redis.dbId"
	RedisMaxIdleConnections = "redis.maxIdleConnections"

	DefaultMaxIdleConnections = 10
)

type RedisConfig struct {
	Host     string
	Port     string
	DbId     int
	PoolSize int
}

func loadRedisConfig() *RedisConfig {
	host := viper.GetString(RedisHost)
	port := viper.GetString(RedisPort)
	dbId := viper.GetInt(RedisDbId)
	maxIdleConnections := loadDefaultConfig(RedisMaxIdleConnections, DefaultMaxIdleConnections)
	return &RedisConfig{
		Host:     host,
		Port:     port,
		DbId:     dbId,
		PoolSize: maxIdleConnections,
	}
}
{% endif %}