package configurator

import "github.com/matthewhartstonge/configurator/diag"

var _ ConfigImplementer = (*ConfigType)(nil)

// ConfigType provides an abstract struct (née abstract base class) to compose
// into concrete config parser implementations.
type ConfigType struct {
	// Config provides
	Config ConfigImplementer
}

// Values implements ConfigParser for returning the underlying parsed values.
func (c *ConfigType) Values() any {
	if c.Config == nil {
		return nil
	}
	return c.Config
}

// Validate expects the implementor to return any errors found in the parsed
// configuration. If any errors are found it aims to provide clear user
// instruction as to where errors were encountered and how one would go about
// fixing it.
// A component is provided so that file configurators can pass through if it is
// global or local configuration.
func (c *ConfigType) Validate(component diag.Component) *diag.Diagnostics {
	return c.Config.Validate(component)
}

// Merge expects the implementor to merge the parsed configuration from the
// concrete implementation and bind it back into the provided domain config.
func (c *ConfigType) Merge(domainConfig any) any {
	return c.Config.Merge(domainConfig)
}
