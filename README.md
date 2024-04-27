<p align="center">
  <img alt="Configurator logo" src="assets/configurator-circle.png" height="150" />
  <h3 align="center">Configurator</h3>
      <p align="center">Opinionated config validator, parser and merger for Go</p>
</p>

---

Configurator provides a pluggable way to parse, validate and merge configuration
for applications across files, environment variables and cli flag arguments.


## Install `configurator`

Install configurator core and then pull in your required parsers:
Due to each parser pulling in 3rd-party dependencies, each is shipped as a
separate module to keep your dependencies as slim as possible.

```shell
# Core
go get github.com/matthewhartstonge/configurator

# Pick and choose:
## File Parsers
go get github.com/matthewhartstonge/configurator/hcl
go get github.com/matthewhartstonge/configurator/json
go get github.com/matthewhartstonge/configurator/toml
go get github.com/matthewhartstonge/configurator/yaml

## Environment Variable Parsers
go get github.com/matthewhartstonge/configurator/envconfig

## Flag Parsers
# todo..
```

## Example

Check out the [example application](/_example) with a full example.

## Documentation

Documentation is hosted at [godoc](https://pkg.go.dev/github.com/matthewhartstonge/configurator)

## Usage

```go
// Define your own ExampleFileConfig{}, ExampleEnvConfig{} and ExampleFlagConfig{}
import (
    "fmt"
    "time"
    
    "github.com/matthewhartstonge/configurator"
    "github.com/matthewhartstonge/configurator/envconfig"
    "github.com/matthewhartstonge/configurator/hcl"
    "github.com/matthewhartstonge/configurator/json"
    "github.com/matthewhartstonge/configurator/toml"
    "github.com/matthewhartstonge/configurator/yaml"
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
		Flag: nil, // TODO: implement flag parsers
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

- [ ] CLI Flag parsing
- [ ] Structured diagnotic error reporting
- [ ] Full documentation for developer happiness
