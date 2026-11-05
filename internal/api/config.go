package api

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServiceName string `envconfig:"APP_NAME" default:"bookmark-service"`
	AppPort     int    `envconfig:"APP_PORT" default:"8080"`
	InstanceID  string `envconfig:"INSTANCE_ID" default:""`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}