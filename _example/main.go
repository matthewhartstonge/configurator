package main

import (
	"fmt"
	"time"

	"github.com/matthewhartstonge/configurator"
	"github.com/matthewhartstonge/configurator/env/envconfig"
	"github.com/matthewhartstonge/configurator/file/hcl"
	"github.com/matthewhartstonge/configurator/file/json"
	"github.com/matthewhartstonge/configurator/file/toml"
	"github.com/matthewhartstonge/configurator/file/yaml"
)

type DomainConfig struct {
	Name            string
	Port            uint16
	BackupFrequency time.Duration
}

func main() {
	var defaults = &DomainConfig{
		Name:            "Default Name",
		Port:            9090,
		BackupFrequency: 24 * time.Hour,
	}

	cfg := &configurator.Config{
		AppName: "ExampleApp",
		Domain:  defaults,
		File: []configurator.ConfigTypeable{
			yaml.New(&ExampleFileConfig{}),
			toml.New(&ExampleFileConfig{}),
			json.New(&ExampleFileConfig{}),
			hcl.New(&ExampleFileConfig{}),
		},
		Env:  envconfig.New(&ExampleEnvConfig{}),
		Flag: nil,
	}

	begin := time.Now()

	appConfig, err := configurator.New(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("took: ", time.Since(begin))

	fmt.Printf("%+v\n", appConfig.Domain)
}
