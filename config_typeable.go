package configurator

import (
	"github.com/matthewhartstonge/configurator/diag"
)

type ConfigTypeable interface {
	ConfigParser
	ConfigImplementer
}

type ConfigFileParser interface {
	// Stat returns false if a file can't be found by the parser.
	Stat(diags *diag.Diagnostics, component diag.Component, cfg *Config, dirPath string) bool
}

type ConfigParser interface {
	// Type informs the user as to which parser is being used.
	Type() string
	// Parse returns the direct file path of the file that was parsed and any
	// associated errors returned from parsing the file.
	Parse(cfg *Config) (string, error)
	// Values returns the current state of the configuration values.
	Values() any
}

type ConfigImplementer interface {
	Validate(component diag.Component) diag.Diagnostics
	Merge(config any) any
}
