package configurator

import (
	"errors"
	"os"
	"path/filepath"
)

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
}

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
func New(config *Config) (*Config, error) {
	diags := []error{}

	// default filename to 'config' if not provided.
	if config.FileName == "" {
		config.FileName = "config"
	}

	// Process OS application directory configuration files.
	globalAppPath, err := getGlobalPath(config)
	if err != nil {
		return config, err
	}
	config, diags = processFiles(config, diags, globalAppPath)

	// Process current working directory configuration files.
	cwd, err := os.Getwd()
	if err != nil {
		return config, err
	}
	config, diags = processFiles(config, diags, cwd)

	// Process environment variable configuration.
	config, diags = processConfig(config, config.Env, diags)

	// TODO: flags config provider
	// Process CLI provided flag arguments.
	config, diags = processConfig(config, config.Flag, diags)

	return config, consolidateErrors(diags)
}

// getGlobalPath returns the OS global path to the application's directory.
func getGlobalPath(cfg *Config) (string, error) {
	// obtain os specific user configuration directory
	appDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return appDir + string(filepath.Separator) + cfg.AppName, nil
}

// processFiles iterates through the provided file type parsers, stating the file
func processFiles(cfg *Config, diags []error, filepath string) (*Config, []error) {
	for _, fileConfig := range cfg.File {
		if fp, ok := fileConfig.(FileParser); ok && !fp.Stat(cfg, filepath) {
			// If we can't find the file, skip it.
			continue
		}

		cfg, diags = processConfig(cfg, fileConfig, diags)
	}

	return cfg, diags
}

func processConfig(cfg *Config, configurer ConfigTypeable, diags []error) (*Config, []error) {
	if configurer == nil {
		// no parser provided
		return cfg, diags
	}

	if err := configurer.Parse(cfg); err != nil {
		diags = append(diags, err)
	}

	diags.Append(configurer.Validate(component)...)

	cfg.Domain = configurer.Merge(cfg.Domain)

	return cfg, diags
}

// consolidateErrors takes a variadic number of errors and pretty prints them
// into a single error.
func consolidateErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	errMsg := ""
	for _, e := range errs {
		if e != nil {
			errMsg += e.Error() + "\n"
		}
	}
	return errors.New(errMsg)
}
