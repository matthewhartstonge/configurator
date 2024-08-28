package main

import (
	"strconv"
	"time"

	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/diag"
)

var _ configurator.ConfigImplementer = (*ExampleFileConfig)(nil)

type ExampleFileConfig struct {
	MyApp struct {
		Name            string `hcl:"name,label" json:"name" toml:"Name" yaml:"name"`
		Port            int    `hcl:"port,optional" json:"port" toml:"Port" yaml:"port"`
		BackupFrequency int    `hcl:"backup_frequency" json:"backupFrequency" toml:"BackupFrequency" yaml:"backup_frequency"`
		Version         string `hcl:"version,optional" json:"version" toml:"Version" yaml:"version"`
	} `hcl:"app,block" json:"myapp" toml:"MyApp" yaml:"myapp"`
}

func (e *ExampleFileConfig) Validate(component diag.Component) diag.Diagnostics {
	var diags diag.Diagnostics
	if e.MyApp.Port < 0 || e.MyApp.Port > 65535 {
		diags.FromComponent(component, "myapp.port").
			Error("Unable to parse port",
				"Port must be between 0 and 65535, but instead got "+strconv.Itoa(e.MyApp.Port))
		e.MyApp.Port = 0
	}
	if e.MyApp.BackupFrequency < 0 {
		diags.FromComponent(component, "myapp.backupFrequency").
			Error("Unable to parse backup frequency",
				"Backup frequency should be provided in hours and should be non-negative, but got "+strconv.Itoa(e.MyApp.BackupFrequency))
		e.MyApp.BackupFrequency = 0
	}

	return diags
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
	if e.MyApp.Version != "" {
		cfg.Version = e.MyApp.Version
	}

	return cfg
}
