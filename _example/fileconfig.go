package main

import (
	"fmt"
	"time"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*ExampleFileConfig)(nil)

type ExampleFileConfig struct {
	MyApp struct {
		Name            string `hcl:"name,label" json:"name" toml:"Name" yaml:"name"`
		Port            int    `hcl:"port,optional" json:"port" toml:"Port" yaml:"port"`
		BackupFrequency int    `hcl:"backup_frequency" json:"backupFrequency" toml:"BackupFrequency" yaml:"backup_frequency"`
	} `hcl:"app,block" json:"myapp" toml:"MyApp" yaml:"myapp"`
}

func (e *ExampleFileConfig) Validate() error {
	if e.MyApp.Port < 0 || e.MyApp.Port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}

	return nil
}

func (e *ExampleFileConfig) Merge(d any) any {
	cfg := d.(*DomainConfig)

	if e.MyApp.Name != "" {
		cfg.Name = e.MyApp.Name
	}
	if e.MyApp.Port != 0 {
		cfg.Port = uint16(e.MyApp.Port)
	}
	if e.MyApp.BackupFrequency != 0 {
		cfg.BackupFrequency = time.Duration(e.MyApp.BackupFrequency) * time.Hour
	}

	return cfg
}
