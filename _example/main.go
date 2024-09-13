package main

import (
	stdjson "encoding/json"
	"fmt"
	"time"

	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/env/envconfig"
	"github.com/matthewhartstonge/configurator/file/hcl"
	"github.com/matthewhartstonge/configurator/file/json"
	"github.com/matthewhartstonge/configurator/file/toml"
	"github.com/matthewhartstonge/configurator/file/yaml"
	"github.com/matthewhartstonge/configurator/flag/stdflag"
)

type DomainConfig struct {
	Name            string
	Port            uint16
	BackupFrequency time.Duration
	Version         string
}

func main() {
	var defaults = &DomainConfig{
		Name:            "Default Name",
		Port:            9090,
		BackupFrequency: 24 * time.Hour,
		Version:         "0.0.0",
	}

	cfg := &configurator.Config{
		AppName: "ExampleApp",
		Domain:  defaults,
		File: []configurator.ConfigFileTypeable{
			yaml.New(&ExampleFileConfig{}),
			toml.New(&ExampleFileConfig{}),
			json.New(&ExampleFileConfig{}),
			hcl.New(&ExampleFileConfig{}),
		},
		Env:  envconfig.New(&ExampleEnvConfig{}),
		Flag: stdflag.New(&ExampleFlagConfig{}),
	}

	// Calling `New` implicitly parses the configuration. If you want to
	// re-process configuration at a later time, you can call
	// `appConfig.Parse()`.
	begin := time.Now()
	appConfig, diags := configurator.New(cfg)
	fmt.Println("Config parsing took: ", time.Since(begin))

	// Our diagnostics can be read to see every little step taken, file read and
	// what wasn't able to parse.
	if diags != nil {
		fmt.Printf("\nConfiguration diagnostics:\n%s\n\n", diags)
	}

	// If we want to read the underlying processed values for every file, envvar
	// and cli flag touched, we can access the app config values. This returns
	// an array specifying the component (global file, local file, envionment
	// variable, cli flag), the path the config from where the config was
	// located and the outcomes of configuration values that were found.
	for _, v := range appConfig.Values() {
		fmt.Printf("Parsed %s config at %s with values: %+v\n", v.Component, v.Path, v.Value)
	}

	// Pretty print the merged config to console!
	mergedConfig, _ := stdjson.MarshalIndent(appConfig.Domain, "", "  ")
	fmt.Printf("\nMerged Config:\n%s\n", string(mergedConfig))
}
