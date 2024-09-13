package diag

type Component int

const (
	// componentInvalid represents an undefined component and should not be
	// used.
	componentInvalid Component = iota
	// ComponentGlobalFile states that the diagnostic comes from a global config
	// file.
	ComponentGlobalFile
	// ComponentLocalFile states that the diagnostic comes from a local config
	// file.
	ComponentLocalFile
	// ComponentEnvVar states that the diagnostic comes from an environment
	// variable.
	ComponentEnvVar
	// ComponentFlag states that the diagnostic comes from a cli flag.
	ComponentFlag
	// ComponentFlagFile states that the diagnostic comes from a CLI specified
	// config file.
	ComponentFlagFile
)

func (c Component) String() string {
	switch c {
	case ComponentGlobalFile:
		return "Global Config File"
	case ComponentLocalFile:
		return "Local Config File"
	case ComponentEnvVar:
		return "Environment Variable"
	case ComponentFlag:
		return "CLI Flag"
	case ComponentFlagFile:
		return "CLI Specified Config File"
	default:
		return "Invalid"
	}
}
