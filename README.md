<p align="center">
  <img alt="Configurator logo" src="assets/configurator-circle.png" height="150" />
  <h3 align="center">Configurator</h3>
  <p align="center">Opinionated config validator, parser and merger for Go</p>
</p>

---

Configurator provides a pluggable way to parse, validate and merge configuration
for applications across files, environment variables and cli flag arguments.


## Install `configurator`

```shell
# Core
go get github.com/matthewhartstonge/configurator
```

## Example

Check out the [example application](/_example) with a full example.

## Documentation

Documentation is hosted at [godoc](https://pkg.go.dev/github.com/matthewhartstonge/configurator)

## Usage

Each parser is stored in a separate package due to pulling in 3rd-party 
dependencies. This ensures that your build is kept slim - includes what you need
and nothing that you don't.

```go
package main

// Define your own FileConfig{}, EnvConfig{} and FlagConfig{}

import (
    "fmt"
    
    "github.com/matthewhartstonge/configurator"
    "github.com/matthewhartstonge/configurator/env/envconfig"
    "github.com/matthewhartstonge/configurator/file/hcl"
    "github.com/matthewhartstonge/configurator/file/json"
    "github.com/matthewhartstonge/configurator/file/toml"
    "github.com/matthewhartstonge/configurator/file/yaml"
    "github.com/matthewhartstonge/configurator/flag/stdflag"
)

func main() {
    cfg := &configurator.Config{
		AppName: "ExampleApp",
		Domain:  defaults,
		File: []configurator.ConfigTypeable{
			yaml.New(&FileConfig{}),
			toml.New(&FileConfig{}),
			json.New(&FileConfig{}),
			hcl.New(&FileConfig{}),
		},
		Env:  envconfig.New(&EnvConfig{}),
        Flag: stdflag.New(&FlagConfig{}),
	}
	config, err := configurator.New(cfg)
	if err != nil {
		panic(err)
	}
	
	// Print out the processed config:
	fmt.Printf("%+v\n", config.Domain)
}
```

## Todo

- [ ] Full documentation for developer happiness
  - [ ] What are the interfaces required to be implemented?