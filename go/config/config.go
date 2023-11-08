package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	{% if postgres_enabled %}
	DBConfig    *DatabaseConfig
	{% endif %}
	LogConfig   *LogConfig
	AppConfig   *AppConfig
	{% if redis_enabled %}
	RedisConfig *RedisConfig
	{% endif %}
}

func loadDefaultConfig[T comparable](key string, t T) T {
	val, ok := viper.Get(key).(T)
	if !ok || val == *new(T) {
		return t
	}
	return val
}

func loadConfigMust[T comparable](key string) T {
	val, ok := viper.Get(key).(T)
	if !ok || val == *new(T) {
		panic("config not found: " + key)
	}
	return val
}

// LoadConfig sets up viper and returns config, or panics if config can't be loaded
func LoadConfig() Config {
	switch os.Getenv("ENV") {
	case "local":
		loadLocalConfig()
	default:
		loadStandardConfig()
	}
	return Config{
		{% if postgres_enabled %}
		DBConfig:    loadDBConfig(),
		{% endif %}
		LogConfig:   loadLogConfig(),
		AppConfig:   loadAppConfig(),
		{% if redis_enabled %}
		RedisConfig: loadRedisConfig(),
		{% endif %}
	}
}

func loadStandardConfig() {
	viper.ReadInConfig()
}

func loadLocalConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
