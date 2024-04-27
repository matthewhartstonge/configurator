package toml

import (
	"github.com/pelletier/go-toml/v2"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*TOML)(nil)

func New(config configurator.ConfigImplementer) *TOML {
	return &TOML{
		FileProvider: configurator.NewFileProvider(
			config,
			[]string{"toml"},
			toml.Unmarshal,
		),
	}
}

type TOML struct {
	configurator.FileProvider
}
