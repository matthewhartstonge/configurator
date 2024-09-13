package configurator

import (
	"github.com/matthewhartstonge/configurator/diag"
)

type ConfigFileTypeable interface {
	ConfigParser
	ConfigFileParser
	ConfigImplementer
}

type ConfigFileParser interface {
	// Stat returns false if a file can't be found by the parser.
	Stat(diags *diag.Diagnostics, component diag.Component, cfg *Config, dirPath string) bool
}
