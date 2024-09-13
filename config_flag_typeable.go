package configurator

type ConfigFlagTypeable interface {
	ConfigParser
	ConfigFlagImplementer
}

type ConfigFlagImplementer interface {
	ConfigImplementer

	// Init is called to configure and set up the flag environment.
	Init()
}
