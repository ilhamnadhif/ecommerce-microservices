package config

const (
	ProductService  = "product"
	MerchantService = "merchant"
	CustomerService = "customer"
	CartService     = "cart"
)

type (
	Configuration struct {
		Server  ServerConfig       `yaml:"server"`
		Service map[string]Service `yaml:"service"`
		Jwt     JwtConfig          `yaml:"jwt"`
	}

	Service struct {
		ServiceName string `yaml:"serviceName"`
		HostPort    string `yaml:"hostPort"`
	}

	ServerConfig struct {
		HostPort string `yaml:"hostPort"`
	}

	JwtConfig struct {
		ExpiredIn  int    `yaml:"expiredIn"`
		SigningKey string `yaml:"signingKey"`
	}
)
