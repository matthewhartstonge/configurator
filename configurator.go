package configurator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/matthewhartstonge/configurator/diag"
)

// New parses the config in order of precedence:
//
// 1. Command line.
// 2. Config file that's name is declared on the command line.
// 3. Environment variables
// 3. Local Config File (if exists)
// 3. Global Config File (if exists)
//
// To be clear, this means config files are searched for and read first, then
// environment variables are merged in over the top, then command line flags as
// the highest priority.
func New(config *Config) (*Config, diag.Diagnostics) {
	return config.Parse()
}

type Config struct {
	// AppName defines the application name.
	//
	// For file based pathing, the application name is used to find the
	// application's directory in the global application directory. For example,
	// on linux, if your application is name "example" it will look
	// for configuration in /etc/example
	AppName string
	// FileName overrides the name of the config file to look for.
	// By default, will search for files named "config".
	FileName string
	// Domain is your own domain specific config from which all other
	// configuration types will be merged into. This struct can define its own
	// specific types where each ConfigImplementer can implement the requisite
	// type casting, validation and merging.
	Domain any
	// File provides a list of file configurators to parse, validate and merge
	// global, current working directory and flag specified config files.
	// Filetypes are processed and merged in specified order. This means that
	// the last filetype specified takes highest precedence in merging.
	File []ConfigTypeable
	// Env provides a configurator to parse, validate and merge configuration
	// variables from the user's environment.
	Env ConfigTypeable
	// Flag provides a configurator to parse, validate and merge configuration
	// variables from the user's specified cli flag arguments.
	Flag ConfigTypeable

	// parsed stores the parsed values of each config.
	parsed []ParsedConfig
}

// ParsedConfig stores the parsed configuration values.
type ParsedConfig struct {
	// Component specifies from where the config values came from.
	Component diag.Component
	// Path specifies either the file path, of environment variable prefix the
	// values came from.
	Path string
	// Value holds the processed values.
	Value any
}

// Parse processes the
func (c *Config) Parse() (*Config, diag.Diagnostics) {
	var diags diag.Diagnostics

	// default filename to 'config' if not provided.
	if c.FileName == "" {
		c.FileName = "config"
	}
	c.parsed = nil

	// Process OS application directory configuration files.
	c, diags = processFileConfig(diags, diag.ComponentGlobalFile, c)

	// Process current working directory configuration files.
	c, diags = processFileConfig(diags, diag.ComponentLocalFile, c)

	// Process environment variable configuration.
	c, diags = processConfig(diags, diag.ComponentEnvVar, c, c.Env)

	// TODO: flags config provider
	// Process CLI provided flag arguments.
	c, diags = processConfig(diags, diag.ComponentFlag, c, c.Flag)

	return c, diags
}

// processFileConfig iterates through the provided file type parsers, stating the file.
func processFileConfig(diags diag.Diagnostics, component diag.Component, cfg *Config) (*Config, diag.Diagnostics) {
	paths, diags := getConfigPaths(diags, component, cfg)

	for _, path := range paths {
		for _, fileConfig := range cfg.File {
			if fp, ok := fileConfig.(ConfigFileParser); ok && !fp.Stat(&diags, component, cfg, path) {
				// If we can't find the file, skip it.
				continue
			}

			cfg, diags = processConfig(diags, component, cfg, fileConfig)
		}
	}

	return cfg, diags
}

// getConfigPaths returns file paths to the configuration directory.
func getConfigPaths(diags diag.Diagnostics, component diag.Component, cfg *Config) ([]string, diag.Diagnostics) {
	var paths []string
	switch component {
	case diag.ComponentGlobalFile:
		if runtime.GOOS == "linux" {
			// Search at /etc/{APP_NAME}
			dir := string(filepath.Separator) + "etc"
			fp := configFP(cfg, dir)
			diags.GlobalFile(dir).Trace("User Configuration Directory Added", fp)
			paths = append(paths, fp)
		}

		if dir, err := os.UserConfigDir(); err != nil {
			diags.GlobalFile(dir).Trace(
				"Unable to Obtain Path to User Configuration Directory",
				fmt.Sprintf("Unable to find path to global configuration '%s' file as %s", cfg.FileName, err.Error()),
			)
		} else {
			fp := configFP(cfg, dir)
			diags.GlobalFile(dir).Trace("User Configuration Directory Added", fp)
			paths = append(paths, fp)
		}

	case diag.ComponentLocalFile:
		if dir, err := os.UserHomeDir(); err != nil {
			diags.LocalFile(dir).Trace(
				"Unable to Obtain Path to User Home Directory",
				fmt.Sprintf("Unable to find path to local configuration '%s' file as %s", cfg.FileName, err.Error()),
			)
		} else {
			fp := configFP(cfg, dir)
			diags.LocalFile(dir).Trace("User Home Directory Added", fp)
			paths = append(paths, fp)
		}

		if dir, err := os.Getwd(); err != nil {
			diags.LocalFile(dir).Trace(
				"Unable to Obtain Path to Current Working Directory",
				fmt.Sprintf("Unable to find path to local configuration '%s' file as %s", cfg.FileName, err.Error()),
			)
		} else {
			// check for a config file directly in the working directory.
			diags.LocalFile(dir).Trace("Current Working Directory Added", dir)
			paths = append(paths, dir)
		}

	default:
		return paths, diags.
			FromComponent(component, "").
			Error("Unknown File Component Supplied",
				fmt.Sprintf(
					"File component %s was supplied, but required either a global or local file. "+
						"This generally indicates a bug in the config parsing provider and should be reported as a bug",
					component,
				),
			)
	}

	return paths, diags
}

// configFP returns a well-formed path to an expected application directory.
func configFP(cfg *Config, dir string) string {
	return dir + string(filepath.Separator) + cfg.AppName
}

// processConfig does the heavy lifting of parsing, validating and merging the
// config together returning diagnostic information at the end of the process.
func processConfig(diags diag.Diagnostics, component diag.Component, cfg *Config, configurer ConfigTypeable) (*Config, diag.Diagnostics) {
	if configurer == nil {
		// no parser provided, may be expected, for example, if CLI flags aren't implemented.
		diags.FromComponent(component, "").
			Trace("No configurator provided",
				fmt.Sprintf("Error attempting to parse %s configuration", component))
		return cfg, diags
	}

	path, err := configurer.Parse(cfg)
	if err != nil {
		// Low-level parsing issue
		diags.FromComponent(component, configurer.String()).
			Error(fmt.Sprintf("Error parsing %s configuration", component),
				err.Error())
		return cfg, diags
	}

	cfg.appendParsedConfig(component, path, configurer.Values())

	diags.Append(configurer.Validate(component)...)

	cfg.Domain = configurer.Merge(cfg.Domain)

	return cfg, diags
}

// appendParsedConfig injects parsed config values for later perusal.
func (c *Config) appendParsedConfig(component diag.Component, path string, v any) {
	c.parsed = append(c.parsed, ParsedConfig{component, path, v})
}

// Values returns the evaluated configuration values.
func (c *Config) Values() []ParsedConfig {
	return c.parsed
}
