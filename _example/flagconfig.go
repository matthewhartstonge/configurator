package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/diag"
	"github.com/matthewhartstonge/configurator/flag/stdflag"
)

var _ configurator.ConfigParser = (*ExampleFlagConfig)(nil)
var _ configurator.ConfigImplementer = (*ExampleFlagConfig)(nil)
var _ configurator.ConfigFlagImplementer = (*ExampleFlagConfig)(nil)

type ExampleFlagConfig struct {
	stdflag.Flag

	Port            int
	BackupFrequency int
}

func (f *ExampleFlagConfig) Init() {
	flag.IntVar(&f.Port, "port", 0, "path to config file")
	flag.IntVar(&f.BackupFrequency, "backup-frequency", 0, "path to config file")
}

func (f *ExampleFlagConfig) Validate(component diag.Component) *diag.Diagnostics {
	diags := new(diag.Diagnostics)
	if f.Port < 0 || f.Port > 65535 {
		diags.FromComponent(component, "-port").
			Error("Unable to parse port",
				"Port must be between 0 and 65535, but instead got "+strconv.Itoa(f.Port))
		f.Port = 0
	}
	if f.BackupFrequency < 0 {
		diags.FromComponent(component, "-backup-frequency").
			Error("Unable to parse backup frequency",
				"Backup frequency should be provided in hours and be non-negative, but instead got "+strconv.Itoa(f.BackupFrequency))
		f.BackupFrequency = 0
	}

	return diags
}

func (f *ExampleFlagConfig) Merge(d any) any {
	cfg := d.(*DomainConfig)

	if f.Port != 0 {
		cfg.Port = uint16(f.Port)
	}
	if f.BackupFrequency != 0 {
		cfg.BackupFrequency = time.Duration(f.BackupFrequency) * time.Hour
	}

	return cfg
}
