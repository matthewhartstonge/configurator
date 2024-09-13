package diag

// Diagnostics contains a report of diagnostic information.
type Diagnostics []Diagnostic

// Append adds a number of new diagnostic entries to the diagnostics.
func (d *Diagnostics) Append(diags ...Diagnostic) {
	if len(diags) == 0 {
		// Nothing to append!
		return
	}

	*d = append(*d, diags...)
}

// GlobalFile enables building up a diagnostic message for a global
// configuration file value.
func (d *Diagnostics) GlobalFile(path string) *Builder {
	return d.builder(ComponentGlobalFile, path)
}

// Env enables building up a diagnostic message for an environment variable value.
func (d *Diagnostics) Env(path string) *Builder {
	return d.builder(ComponentEnvVar, path)
}

// LocalFile enables building up a diagnostic message for a local configuration
// file value.
func (d *Diagnostics) LocalFile(path string) *Builder {
	return d.builder(ComponentLocalFile, path)
}

// Flag enables building up a diagnostic message for a CLI flag value.
func (d *Diagnostics) Flag(path string) *Builder {
	return d.builder(ComponentFlag, path)
}

// FlagFile enables building up a diagnostic message for a CLI specified config
// file value.
func (d *Diagnostics) FlagFile(path string) *Builder {
	return d.builder(ComponentFlagFile, path)
}

// FromComponent enables taking in a component enum to build up a diagnostic
// message.
func (d *Diagnostics) FromComponent(component Component, path string) *Builder {
	return d.builder(component, path)
}

// builder returns a diagnostic builder API for creating a diagnostic message.
func (d *Diagnostics) builder(component Component, path string) *Builder {
	if d == nil {
		d = &Diagnostics{}
	}
	return &Builder{
		d: d,
		e: &Diagnostic{Component: component, Path: path},
	}
}

// Fatals returns all diagnostic entries at SeverityFatal level.
func (d *Diagnostics) Fatals() Diagnostics {
	return d.getDiagsWithLevel(SeverityFatal)
}

// Errors returns all diagnostic entries at SeverityError level.
func (d *Diagnostics) Errors() Diagnostics {
	return d.getDiagsWithLevel(SeverityError)
}

// Warnings returns all diagnostic entries at SeverityWarn level.
func (d *Diagnostics) Warnings() Diagnostics {
	return d.getDiagsWithLevel(SeverityWarn)
}

// Infos returns all diagnostic entries at SeverityInfo level.
func (d *Diagnostics) Infos() Diagnostics {
	return d.getDiagsWithLevel(SeverityInfo)
}

// Debugs returns all diagnostic entries at SeverityDebug level.
func (d *Diagnostics) Debugs() Diagnostics {
	return d.getDiagsWithLevel(SeverityDebug)
}

// Traces returns all diagnostic entries at SeverityTrace level.
func (d *Diagnostics) Traces() Diagnostics {
	return d.getDiagsWithLevel(SeverityTrace)
}

// getDiagsWithLevel returns an array of diagnostics that match the specified
// severity level.
func (d Diagnostics) getDiagsWithLevel(sev Severity) Diagnostics {
	var diags Diagnostics
	for _, diag := range d {
		if diag.Severity != sev {
			continue
		}

		diags = append(diags, diag)
	}

	return diags
}
