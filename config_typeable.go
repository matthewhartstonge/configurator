package configurator

import "github.com/matthewhartstonge/configurator/diag"

type ConfigTypeable interface {
	ConfigParser
	ConfigImplementer
}

type ConfigFileParser interface {
	Stat(diags *diag.Diagnostics, component diag.Component, cfg *Config, dirPath string) bool
}

type ConfigParser interface {
	Parse(cfg *Config) error
}

type ConfigImplementer interface {
	Validate(component diag.Component) diag.Diagnostics
	Merge(config any) any
}
