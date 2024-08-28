package stdflag

import (
	"flag"
	"os"

	"github.com/matthewhartstonge/configurator"
)

var (
	_ configurator.ConfigParser          = (*Flag)(nil)
	_ configurator.ConfigFlagImplementer = (*Flag)(nil)
	_ configurator.ConfigImplementer     = (*Flag)(nil)
)

func New(config configurator.ConfigFlagImplementer) *Flag {
	return &Flag{
		ConfigType: configurator.ConfigType{
			Config: config,
		},
	}
}

type Flag struct {
	configurator.ConfigType
}

func (f *Flag) Init() {
	if f == nil || f.Config == nil {
		return
	}

	if f.Config != nil {
		f.Config.(configurator.ConfigFlagImplementer).Init()
	}
}

func (f *Flag) Type() string {
	return "stdflag configurator"
}

func (f *Flag) Parse(_ *configurator.Config) (string, error) {
	return "args", flag.CommandLine.Parse(os.Args[1:])
}
