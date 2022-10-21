package config

const (
	MerchantService = "merchant"
	CustomerService = "customer"
)

type (
	Configuration struct {
		Server  ServerConfig       `yaml:"server"`
		Service map[string]Service `yaml:"service"`
		Redis   RedisConfig        `yaml:"redis"`
		Jwt     JwtConfig          `yaml:"jwt"`
	}

	ServerConfig struct {
		ServiceName string `yaml:"serviceName"`
		HostPort    string `yaml:"hostPort"`
	}

	Service struct {
		ServiceName string `yaml:"serviceName"`
		HostPort    string `yaml:"hostPort"`
	}

	RedisConfig struct {
		HostPort string `yaml:"hostPort"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbNumber int    `yaml:"dbNumber"`
	}
	JwtConfig struct {
		ExpiredIn  int    `yaml:"expiredIn"`
		SigningKey string `yaml:"signingKey"`
	}
)
