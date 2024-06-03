package main

import (
	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/diag"
)

var _ configurator.ConfigImplementer = (*ExampleEnvConfig)(nil)

type ExampleEnvConfig struct {
	Name string `envconfig:"NAME"`
	Port int    `envconfig:"PORT" default:"9090"`
}

func (e *ExampleEnvConfig) Validate(_ diag.Component) diag.Diagnostics {
	var diags diag.Diagnostics
	if e.Port < 0 || e.Port > 65535 {
		diags.Env("PORT").Error("Unable to parse port", "Port must be between 0 and 65535")
	}

	return diags
}

func (e *ExampleEnvConfig) Merge(d any) any {
	cfg := d.(*DomainConfig)

	if e.Name != "" {
		cfg.Name = e.Name
	}

	if e.Port != 0 {
		cfg.Port = uint16(e.Port)
	}

	return cfg
}
