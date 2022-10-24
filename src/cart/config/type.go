package config

type (
	Configuration struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		Redis    RedisConfig    `yaml:"redis"`
	}

	ServerConfig struct {
		ServiceName string `yaml:"serviceName"`
		HostPort    string `yaml:"hostPort"`
	}

	DatabaseConfig struct {
		Driver   string `yaml:"driver"`
		HostPort string `yaml:"hostPort"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}

	RedisConfig struct {
		HostPort string `yaml:"hostPort"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbNumber int    `yaml:"dbNumber"`
	}
)
