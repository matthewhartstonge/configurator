package toml

import (
	toml "github.com/pelletier/go-toml/v2"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*TOML)(nil)

func New(config configurator.ConfigImplementer) *TOML {
	return &TOML{
		ConfigFileType: configurator.NewConfigFileType(
			config,
			[]string{"toml"},
			toml.Unmarshal,
		),
	}
}

type TOML struct {
	configurator.ConfigFileType
}

func (t TOML) Type() string {
	return "TOML configurator"
}
