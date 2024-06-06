package main

import (
	"flag"
	"time"

	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/diag"
)

var _ configurator.ConfigImplementer = (*ExampleFileConfig)(nil)

func init() { // TODO: plug this into configurator via a method
	flag.StringVar(&flgCfg.ConfigPath, "config", "", "Provides a path to a specific config file")
	flag.IntVar(&flgCfg.Port, "port", 0, "Defines the server port")
	flag.IntVar(&flgCfg.BackupFrequency, "backup-frequency", 0, "Defines the frequency of backups in hours")
}

var (
	flgCfg = &ExampleFlagConfig{} // TODO: plug this into configurator via a method
)

type ExampleFlagConfig struct {
	ConfigPath      string
	Port            int
	BackupFrequency int
}

func (e *ExampleFlagConfig) Validate(component diag.Component) diag.Diagnostics {
	var diags diag.Diagnostics
	if e.Port < 0 || e.Port > 65535 {
		diags.FromComponent(component, "--port").
			Error("Unable to parse port",
				"Port must be between 0 and 65535")
	}
	if e.BackupFrequency < 0 {
		diags.FromComponent(component, "--backup-frequency").
			Error("Unable to parse backup frequency",
				"Backup frequency should be provided in hours")
	}

	return nil
}

func (e *ExampleFlagConfig) Merge(d any) any {
	cfg := d.(*DomainConfig)

	if e.Port != 0 {
		cfg.Port = uint16(e.Port)
	}
	if e.BackupFrequency != 0 {
		cfg.BackupFrequency = time.Duration(e.BackupFrequency) * time.Hour
	}

	return cfg
}
