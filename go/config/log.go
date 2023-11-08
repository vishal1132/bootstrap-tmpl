package config

type LogConfig struct {
	Level     string
	Namespace string
}

func loadLogConfig() *LogConfig {
	return &LogConfig{
		Level:     loadDefaultConfig[string]("log.level", "INFO"),
		Namespace: loadConfigMust[string]("log.namespace"),
	}
}
