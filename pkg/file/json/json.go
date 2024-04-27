package json

import (
	"encoding/json"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigImplementer = (*JSON)(nil)

func New(config configurator.ConfigImplementer) *JSON {
	return &JSON{
		FileProvider: configurator.NewFileProvider(
			config,
			[]string{"json"},
			json.Unmarshal,
		),
	}
}

type JSON struct {
	configurator.FileProvider
}
