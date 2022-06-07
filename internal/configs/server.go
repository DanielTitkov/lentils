package configs

import "fmt"

type ServerConfig struct {
	Port    int
	TLSPort int `yaml:"tlsPort"`
	Host    string
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
