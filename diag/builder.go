package diag

// Builder provides a log entry building api. This should not be used directly,
// but should be used by calling the component based method on the supplied
// Diagnostics.
type Builder struct {
	d *Diagnostics
	e *Diagnostic
}

func (b *Builder) Fatal(summary, detail string) *Diagnostics {
	return b.build(SeverityFatal, summary, detail)
}

func (b *Builder) Error(summary, detail string) *Diagnostics {
	return b.build(SeverityError, summary, detail)
}

func (b *Builder) Warn(summary, detail string) *Diagnostics {
	return b.build(SeverityWarn, summary, detail)
}

func (b *Builder) Info(summary, detail string) *Diagnostics {
	return b.build(SeverityInfo, summary, detail)
}

func (b *Builder) Debug(summary, detail string) *Diagnostics {
	return b.build(SeverityDebug, summary, detail)
}

func (b *Builder) Trace(summary, detail string) *Diagnostics {
	return b.build(SeverityTrace, summary, detail)
}

func (b *Builder) build(sev Severity, summary, detail string) *Diagnostics {
	diags, diag := b.d, b.e

	diag.Severity = sev
	diag.Summary = summary
	diag.Detail = detail

	diags.Append(*diag)

	return diags
}
