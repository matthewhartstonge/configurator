package diag

// Diagnostics contains a report of diagnostic information.
type Diagnostics struct {
	// diags contains the list of diagnostic entries.
	diags []Diagnostic

	// HasFatal reports if the diagnostics have reported a fatal issue.
	HasFatal bool
	// HasError reports if the diagnostics have reported an error.
	HasError bool
	// HasWarn reports if the diagnostics have reported a warning.
	HasWarn bool
}

// Len reports the number of diagnostic messages logged.
func (d *Diagnostics) Len() int {
	if d == nil {
		return 0
	}

	return len(d.diags)
}

// All returns all diagnostic messages.
func (d *Diagnostics) All() []Diagnostic {
	if d == nil {
		return nil
	}

	return d.diags
}

// Append adds a number of new diagnostic entries to the diagnostics.
func (d *Diagnostics) Append(diags ...Diagnostic) {
	if len(diags) == 0 {
		// Nothing to append!
		return
	}

	for _, diag := range diags {
		switch diag.Severity {
		case SeverityFatal:
			d.HasFatal = true
		case SeverityError:
			d.HasError = true
		case SeverityWarn:
			d.HasWarn = true
		}

		d.diags = append(d.diags, diag)
	}
}

// Merge appends the provided diags into the diagnostics.
func (d *Diagnostics) Merge(diags Diagnostics) {
	if diags.Len() == 0 {
		// Nothing to append!
		return
	}

	if diags.HasFatal {
		d.HasFatal = diags.HasFatal
	}
	if diags.HasError {
		d.HasError = diags.HasError
	}
	if diags.HasWarn {
		d.HasWarn = diags.HasWarn
	}

	d.diags = append(d.diags, diags.All()...)
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
func (d *Diagnostics) getDiagsWithLevel(sev Severity) Diagnostics {
	var diags Diagnostics
	for _, diag := range d.diags {
		if diag.Severity != sev {
			continue
		}

		diags.Append(diag)
	}

	return diags
}
