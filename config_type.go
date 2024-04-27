package configurator

type ConfigType struct {
	Data ConfigImplementer
}

func (c *ConfigType) Validate() error {
	// base case no validation required
	return nil
}

func (c *ConfigType) Merge(config any) any {
	panic("implement me")
}
