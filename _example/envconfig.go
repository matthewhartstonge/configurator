package main

import (
	"fmt"
	
	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*ExampleEnvConfig)(nil)

type ExampleEnvConfig struct {
	Name string `envconfig:"NAME"`
	Port int    `envconfig:"PORT" default:"9090"`
}

func (e *ExampleEnvConfig) Validate() error {
	if e.Port < 0 || e.Port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}

	return nil
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
