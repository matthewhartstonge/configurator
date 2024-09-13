package configurator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/matthewhartstonge/configurator/diag"
)

// New calls parse and returns merged config values in order of precedence:
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
	// FileFlag overrides the flag name used to process a config file at a
	// specified place.
	FileFlag string
	// Domain is your own domain specific config from which all other
	// configuration types will be merged into. This struct can define its own
	// specific types where each ConfigImplementer can implement the requisite
	// type casting, validation and merging.
	Domain any

	// ConfigFilePath stores the path specified via the `-config` CLI flag.
	// If non-empty, configurator will process the config file specified instead
	// of attempting to find global or local config files.
	ConfigFilePath string
	// File provides a list of file configurators to parse, validate and merge
	// global, current working directory and flag specified config files.
	// Filetypes are processed and merged in specified order. This means that
	// the last filetype specified takes highest precedence in merging.
	File []ConfigFileTypeable
	// Env provides a configurator to parse, validate and merge configuration
	// variables from the user's environment.
	Env ConfigTypeable
	// Flag provides a configurator to parse, validate and merge configuration
	// variables from the user's specified cli flag arguments.
	Flag ConfigFlagTypeable

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
		c.FileName = DEFAULT_CONFIG_FILENAME
	}
	c.parsed = nil

	diags = c.processFileFlagConfig(diags)

	if c.ConfigFilePath != "" {
		// Process OS application directory configuration files.
		diags = c.processFileConfig(diags, diag.ComponentFlagFile)
	} else {
		// Process OS application directory configuration files.
		diags = c.processFileConfig(diags, diag.ComponentGlobalFile)

		// Process current working directory configuration files.
		diags = c.processFileConfig(diags, diag.ComponentLocalFile)

		// Process environment variable configuration.
		diags = c.processConfig(diags, diag.ComponentEnvVar, c.Env)
	}

	// Process CLI provided flag configuration.
	diags = c.processFlagConfig(diags, diag.ComponentFlag)

	return c, diags
}

// processFileFlagConfig extracts the path to a config file, if specified via
// the customisable `-config-file` flag.
func (c *Config) processFileFlagConfig(diags diag.Diagnostics) diag.Diagnostics {
	if c.FileFlag == "" {
		c.FileFlag = DEFAULT_CONFIG_FILEFLAG
	}

	// fully-qualified file flag.
	fqFileFlag := "-" + c.FileFlag

	// manually extract the value for the set config file flag.
	v, ok := getFlagValue(c.FileFlag)
	if !ok {
		diags.FlagFile(fqFileFlag).
			Trace("CLI specified config file path not set",
				"Either the value was never set, or an empty string was provided")
		return diags
	}

	c.ConfigFilePath = v
	diags.FlagFile(fqFileFlag).Trace("CLI specified config file path added", c.ConfigFilePath)

	// Remove the flag from os.Args
	removeFlagFromArgs(c.FileFlag)

	return diags
}

// getFlagValue extracts the provided flag name from os.Args manually.
func getFlagValue(name string) (string, bool) {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "-"+name) {
			if strings.Contains(arg, "=") {
				// Handle -flag=value syntax
				return strings.TrimSpace(strings.SplitN(arg, "=", 2)[1]), true
			}
			if i+1 < len(os.Args) {
				// Handle -flag value syntax
				return strings.TrimSpace(os.Args[i+1]), true
			}
		}
	}

	return "", false
}

// removeFlagFromArgs removes the flag and it's value from the global os.Args.
func removeFlagFromArgs(name string) {
	newArgs := make([]string, 0, len(os.Args))

	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "-"+name) {
			if strings.Contains(arg, "=") {
				continue // remove `-flag=value` styled arg
			}
			i++ // remove `-flag value` style styled arg
		}
		newArgs = append(newArgs, arg)
	}

	os.Args = newArgs
}

// processFileConfig iterates through the provided file type parsers, stating the file.
func (c *Config) processFileConfig(diags diag.Diagnostics, component diag.Component) diag.Diagnostics {
	paths, diags := getConfigPaths(diags, component, c)

	for _, path := range paths {
		for _, fileConfig := range c.File {
			if !fileConfig.Stat(&diags, component, c, path) {
				// If we can't find the file, skip it.
				continue
			}

			// process the first found config file based on file type priority.
			return c.processConfig(diags, component, fileConfig)
		}
	}

	return diags
}

// getConfigPaths returns file paths to the configuration directory.
func getConfigPaths(diags diag.Diagnostics, component diag.Component, cfg *Config) ([]string, diag.Diagnostics) {
	if pathStrategy, ok := configFilePathStrategies[component]; ok {
		return pathStrategy(diags, cfg)
	}

	return nil, diags.
		FromComponent(component, "").
		Error("Unknown File Component Supplied",
			fmt.Sprintf(
				"File component %s was supplied, but required either a global or local file. "+
					"This generally indicates a bug in the config parsing provider and should be reported as a bug",
				component,
			),
		)
}

type configFilePathStrategy func(diags diag.Diagnostics, cfg *Config) ([]string, diag.Diagnostics)

var configFilePathStrategies = map[diag.Component]configFilePathStrategy{
	diag.ComponentGlobalFile: processGlobalFilePaths,
	diag.ComponentLocalFile:  processLocalFilePaths,
	diag.ComponentFlagFile:   processFlagFilePath,
}

func processGlobalFilePaths(diags diag.Diagnostics, cfg *Config) ([]string, diag.Diagnostics) {
	var paths []string

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

	return paths, diags
}

func processLocalFilePaths(diags diag.Diagnostics, cfg *Config) ([]string, diag.Diagnostics) {
	var paths []string

	if dir, err := os.UserHomeDir(); err != nil {
		diags.LocalFile(dir).Trace(
			"Unable to obtain path to user home directory",
			fmt.Sprintf("Unable to find path to local configuration '%s' file as %s", cfg.FileName, err.Error()),
		)
	} else {
		fp := configFP(cfg, dir)
		diags.LocalFile(dir).Trace("User home directory added", fp)
		paths = append(paths, fp)
	}

	if dir, err := os.Getwd(); err != nil {
		diags.LocalFile(dir).Trace(
			"Unable to obtain path to current working directory",
			fmt.Sprintf("Unable to find path to local configuration '%s' file as %s", cfg.FileName, err.Error()),
		)
	} else {
		// check for a config file directly in the working directory.
		diags.LocalFile(dir).Trace("Current working directory added", dir)
		paths = append(paths, dir)
	}

	return paths, diags
}

func processFlagFilePath(diags diag.Diagnostics, cfg *Config) ([]string, diag.Diagnostics) {
	if cfg.ConfigFilePath == "" {
		return []string{}, diags
	}

	fqFileFlag := "-" + cfg.FileFlag

	fp := cfg.ConfigFilePath
	absFP, err := filepath.Abs(fp)
	if err != nil {
		diags.FlagFile(fqFileFlag).Error("Unable to compute the absolute file path", err.Error())
		return []string{}, diags
	}

	diags.FlagFile(fqFileFlag).Trace("CLI specified config file path added", absFP)
	return []string{absFP}, diags
}

// configFP returns a well-formed path to an expected application directory.
func configFP(cfg *Config, dir string) string {
	return dir + string(filepath.Separator) + cfg.AppName
}

// processFlagConfig processes and merges in any provided flag configuration.
func (c *Config) processFlagConfig(diags diag.Diagnostics, component diag.Component) diag.Diagnostics {
	if c.Flag == nil {
		return diags
	}

	c.Flag.Init()

	return c.processConfig(diags, component, c.Flag)
}

// processConfig does the heavy lifting of parsing, validating and merging the
// config together returning diagnostic information at the end of the process.
func (c *Config) processConfig(diags diag.Diagnostics, component diag.Component, configurer ConfigTypeable) diag.Diagnostics {
	if configurer == nil {
		// no parser provided, may be expected, for example, if CLI flags aren't implemented.
		diags.FromComponent(component, "").
			Trace("No configurator provided",
				fmt.Sprintf("Error attempting to parse %s configuration", component))
		return diags
	}

	path, err := configurer.Parse(c)
	if err != nil {
		// Low-level parsing issue
		diags.FromComponent(component, configurer.Type()).
			Error(fmt.Sprintf("Error parsing %s configuration", component),
				err.Error())
		return diags
	}

	c.appendParsedConfig(component, path, configurer.Values())

	diags.Append(configurer.Validate(component)...)

	c.Domain = configurer.Merge(c.Domain)

	return diags
}

// appendParsedConfig injects parsed config values for later perusal.
func (c *Config) appendParsedConfig(component diag.Component, path string, v any) {
	c.parsed = append(c.parsed, ParsedConfig{component, path, v})
}

// Values returns the evaluated configuration values.
func (c *Config) Values() []ParsedConfig {
	return c.parsed
}
