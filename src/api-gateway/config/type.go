package config

type (
	Configuration struct {
		Server  ServerConfig       `yaml:"server"`
		Service map[string]Service `yaml:"service"`
	}

	Service struct {
		ServiceName string `yaml:"serviceName"`
		HostPort    string `yaml:"hostPort"`
	}

	ServerConfig struct {
		HostPort string `yaml:"hostPort"`
	}
)
