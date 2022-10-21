package config

type (
	Configuration struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		Redis    RedisConfig    `yaml:"redis"`
		Jwt      JwtConfig      `yaml:"jwt"`
	}

	ServerConfig struct {
		HostPort string `yaml:"hostPort"`
	}

	JwtConfig struct {
		SigningKey string `yaml:"signingKey"`
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
