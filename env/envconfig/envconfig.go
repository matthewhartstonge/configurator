package envconfig

import (
	"strings"

	"github.com/kelseyhightower/envconfig"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*EnvConfig)(nil)

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

func (e EnvConfig) Type() string {
	return "EnvConfig configurator"
}

func (e *EnvConfig) Parse(cfg *configurator.Config) (string, error) {
	return strings.ToTitle(cfg.AppName), envconfig.Process(cfg.AppName, e.Config)
}
