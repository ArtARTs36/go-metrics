package metrics

import "time"

type Config struct {
	Server    ServerConfig `envPrefix:"SERVER_"`
	Namespace string       `env:"NAMESPACE,required"`
}

type ServerConfig struct {
	Addr    string        `env:"ADDR,required"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"30s"`
}
