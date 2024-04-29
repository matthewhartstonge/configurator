package configurator

import (
	"os"
	"path/filepath"
)

// NewConfigFileType provides most functionality required to support a new file
// type. If manual unmarshaling is required, the Unmarshaler can be provided.
func NewConfigFileType(
	config ConfigImplementer,
	fileExtensions []string,
	unmarshaler Unmarshaler,
) ConfigFileType {
	return ConfigFileType{
		unmarshaler: unmarshaler,
		Extensions:  fileExtensions,
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
	// Extensions is a list of file extensions that the provider will look for.
	Extensions []string

	// ConfigType is the embedded configurator.ConfigType.
	ConfigType
}

// Stat checks if the file exists and computes the platform specific Path.
func (f *ConfigFileType) Stat(cfg *Config, dirPath string) bool {
	for _, ext := range f.Extensions {
		cfgFilePath := dirPath + string(filepath.Separator) + cfg.FileName + "." + ext
		if _, err := os.Stat(cfgFilePath); err == nil {
			f.Path = cfgFilePath
			return true
		}
	}
	return false
}

// Parse reads the file based on the generated path computed from Stat and
// unmarshals it into the Config field.
func (f *ConfigFileType) Parse(_ *Config) error {
	file, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}

	return f.unmarshaler(file, f.Config)
}
