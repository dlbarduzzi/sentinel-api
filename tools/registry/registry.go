package registry

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultConfigPath = "."
	defaultConfigType = "env"
	defaultConfigName = ".env"
)

type Config struct {
	Path      string
	Type      string
	Name      string
	EnvPrefix string
}

func New() (*viper.Viper, error) {
	return NewWithConfig(Config{
		Path: defaultConfigPath,
		Type: defaultConfigType,
		Name: defaultConfigName,
	})
}

func NewWithConfig(config Config) (*viper.Viper, error) {
	if config.Path == "" {
		config.Path = defaultConfigPath
	}

	if config.Type == "" {
		config.Type = defaultConfigType
	}

	if config.Name == "" {
		config.Name = defaultConfigName
	}

	v := viper.New()
	v.AutomaticEnv()

	prefix := strings.TrimSpace(config.EnvPrefix)
	if prefix != "" {
		v.SetEnvPrefix(prefix)
	}

	v.AddConfigPath(config.Path)
	v.SetConfigType(config.Type)
	v.SetConfigName(config.Name)

	if err := v.ReadInConfig(); err != nil {
		return v, err
	}

	return v, nil
}
