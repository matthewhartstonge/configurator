package diags

// Diagnostics contains a report of diagnostic information.
type Diagnostics []Diagnostic

// Append adds a number of new diagnostic entries to the diagnostics.
func (d *Diagnostics) Append(diags ...Diagnostic) {
	*d = append(*d, diags...)
}

// Fatals returns all diagnostic entries at SeverityFatal level.
func (d *Diagnostics) Fatals() []Diagnostic {
	return d.getDiagsWithLevel(SeverityFatal)
}

// Errors returns all diagnostic entries at SeverityError level.
func (d *Diagnostics) Errors() []Diagnostic {
	return d.getDiagsWithLevel(SeverityError)
}

// Warnings returns all diagnostic entries at SeverityWarn level.
func (d *Diagnostics) Warnings() []Diagnostic {
	return d.getDiagsWithLevel(SeverityWarn)
}

// Infos returns all diagnostic entries at SeverityInfo level.
func (d *Diagnostics) Infos() []Diagnostic {
	return d.getDiagsWithLevel(SeverityInfo)
}

// Debugs returns all diagnostic entries at SeverityDebug level.
func (d *Diagnostics) Debugs() []Diagnostic {
	return d.getDiagsWithLevel(SeverityDebug)
}

// Traces returns all diagnostic entries at SeverityTrace level.
func (d *Diagnostics) Traces() []Diagnostic {
	return d.getDiagsWithLevel(SeverityTrace)
}

// getDiagsWithLevel returns an array of diagnostics that match the specified
// severity level.
func (d Diagnostics) getDiagsWithLevel(sev Severity) []Diagnostic {
	var diags Diagnostics
	for _, diag := range d {
		if diag.Severity != sev {
			continue
		}

		diags = append(diags, diag)
	}

	return diags
}
