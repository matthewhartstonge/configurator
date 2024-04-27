package configurator

type ConfigTypeable interface {
	ConfigParser
	ConfigImplementer
}

type FileParser interface {
	Stat(cfg *Config, dirPath string) bool
}

type ConfigParser interface {
	Parse(cfg *Config) error
}

type ConfigImplementer interface {
	Validate() error
	Merge(config any) any
}