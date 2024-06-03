package json

import (
	"encoding/json"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*JSON)(nil)

func New(config configurator.ConfigImplementer) *JSON {
	return &JSON{
		ConfigFileType: configurator.NewConfigFileType(
			config,
			[]string{"json"},
			json.Unmarshal,
		),
	}
}

type JSON struct {
	configurator.ConfigFileType
}

func (j JSON) String() string {
	return "JSON configurator"
}
