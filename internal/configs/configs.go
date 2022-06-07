package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env      string
	DB       DBConfig
	Auth     AuthConfig
	Server   ServerConfig
	Data     DataConfig
	External ExternalConfig
	App      AppConfig
	Debug    DebugConfig
}

func ReadConfigs(path string) (Config, error) {
	var cfg Config
	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
