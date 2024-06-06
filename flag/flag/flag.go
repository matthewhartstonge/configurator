package flag

import (
	"flag"
	"os"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*Flag)(nil)

func New(config configurator.ConfigImplementer) *Flag {
	return &Flag{
		ConfigType: configurator.ConfigType{
			Config: config,
		},
	}
}

type Flag struct {
	configurator.ConfigType
}

func (e Flag) Type() string {
	return "Flag configurator"
}

func (e *Flag) Parse(_ *configurator.Config) (string, error) {
	return "args", flag.CommandLine.Parse(os.Args[1:])
}
