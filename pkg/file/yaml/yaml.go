package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*YAML)(nil)

func New(config configurator.ConfigImplementer) *YAML {
	return &YAML{
		FileProvider: configurator.NewFileProvider(
			config,
			[]string{"yaml", "yml"},
			yaml.Unmarshal,
		),
	}
}

type YAML struct {
	configurator.FileProvider
}
