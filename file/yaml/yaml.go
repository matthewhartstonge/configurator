package yaml

import (
	yaml "gopkg.in/yaml.v3"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*YAML)(nil)

func New(config configurator.ConfigImplementer) *YAML {
	return &YAML{
		ConfigFileType: configurator.NewConfigFileType(
			config,
			[]string{"yaml", "yml"},
			yaml.Unmarshal,
		),
	}
}

type YAML struct {
	configurator.ConfigFileType
}

func (y YAML) String() string {
	return "YAML configurator"
}
