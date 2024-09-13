package configurator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matthewhartstonge/configurator/diag"
)

var (
	_ ConfigParser      = (*ConfigFileType)(nil)
	_ ConfigFileParser  = (*ConfigFileType)(nil)
	_ ConfigImplementer = (*ConfigFileType)(nil)
)

// NewConfigFileType provides most functionality required to support a new file
// type. If manual unmarshaling is required, the Unmarshaler can be provided.
func NewConfigFileType(
	config ConfigImplementer,
	fileTypes []string,
	unmarshaler Unmarshaler,
) ConfigFileType {
	return ConfigFileType{
		unmarshaler: unmarshaler,
		Types:       fileTypes,
		ConfigType: ConfigType{
			Config: config,
		},
	}
}

// Unmarshaler unmarshals a byte slice into the given interface.
type Unmarshaler func(data []byte, v interface{}) error

// ConfigFileType provides a ConfigImplementer that reads a file from disk and
// unmarshals it into the Config field. Unmarshaling expects implementations to
// match the standard library interface for Unmarshal.
type ConfigFileType struct {
	// unmarshaler is a function that unmarshals a byte slice into a given
	// interface.
	unmarshaler Unmarshaler
	// Path is the path to the file that will be read and unmarshaled.
	Path string
	// Types is a list of file types or dot extensions that the provider
	// is able to process.
	Types []string

	// ConfigType is the embedded configurator.ConfigType.
	ConfigType
}

// Type returns which parser is in use.
func (f *ConfigFileType) Type() string {
	return "Not Implemented"
}

// Stat checks if the file exists and computes the platform specific Path and
// directly writes to the provided diagnostics.
func (f *ConfigFileType) Stat(diags *diag.Diagnostics, component diag.Component, cfg *Config, filePath string) bool {
	// todo: tidy `ConfigFileType.Stat` implementation. there should be a better way.
	filename := filepath.Base(filePath)
	fileExt := filepath.Ext(filePath)
	for _, fileType := range f.Types {
		// stat for full paths, if provided.
		if fileExt != "" {
			if fileExt != "."+fileType {
				// full file path provided, ext does match provider file type - skip.
				diags.FromComponent(component, filePath).
					Trace("Skipping File Type",
						"The file type does not match "+fileType)
				continue
			}

			// stat full path!
			info, err := os.Stat(filePath)
			if err != nil {
				diags.FromComponent(component, filePath).
					Trace("Config File Not Found",
						"No config file was found at the specified path, error: "+err.Error())
				return false
			}

			if info.IsDir() {
				// y u disguised as file...
				continue
			}

			// specified config file exists for the given file parser!
			f.Path = filePath
			diags.FromComponent(component, filePath).
				Trace("Config File Found",
					fmt.Sprintf("Will attempt to parse %s", filename))
			return true
		}

		// Dynamically build the expected config file path that can be parsed
		// with this provider to check for files existence.
		cfgFilePath := filePath + string(filepath.Separator) + cfg.FileName + "." + filePath
		if _, err := os.Stat(cfgFilePath); err == nil {
			f.Path = cfgFilePath
			diags.FromComponent(component, filePath).
				Trace("Config File Found",
					fmt.Sprintf("Will attempt to parse %s", cfgFilePath))
			return true
		}
	}

	diags.FromComponent(component, filePath).
		Trace("Config File Not Found",
			fmt.Sprintf("Unable to find config file for extensions {%s} at %s", strings.Join(f.Types, ", "), filePath))
	return false
}

// Parse reads the file based on the generated path computed from Stat and
// unmarshals it into the Config field.
func (f *ConfigFileType) Parse(_ *Config) (string, error) {
	file, err := os.ReadFile(f.Path)
	if err != nil {
		return f.Path, err
	}

	return f.Path, f.unmarshaler(file, f.Config)
}
