package envconfig

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*EnvConfig)(nil)

func New(config configurator.ConfigImplementer) *EnvConfig {
	return &EnvConfig{
		ConfigType: configurator.ConfigType{
			Config: config,
		},
	}
}

type EnvConfig struct {
	configurator.ConfigType
}

func (e *EnvConfig) Parse(cfg *configurator.Config) error {
	return envconfig.Process(cfg.AppName, e.Config)
}
