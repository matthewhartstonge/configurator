package configurator

var _ ConfigImplementer = (*ConfigType)(nil)

// ConfigType provides an abstract struct (née abstract base class) to compose
// into concrete config parser implementations.
type ConfigType struct {
	// Config provides
	Config ConfigImplementer
}

// Validate expects the implementor to return any errors found in the parsed
// configuration. If any errors are found it aims to provide clear user
// instruction as to where errors were encountered and how one would go about
// fixing it.
func (c *ConfigType) Validate() error {
	return c.Config.Validate()
}

// Merge expects the implementor to merge the parsed configuration from the
// concrete implementation and bind it back into the provided domain config.
func (c *ConfigType) Merge(domainConfig any) any {
	return c.Config.Merge(domainConfig)
}
