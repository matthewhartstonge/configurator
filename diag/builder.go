package diag

// Builder provides a log entry building api. This should not be used directly,
// but should be used by calling the component based method on the supplied
// Diagnostics.
type Builder struct {
	d *Diagnostics
	e *Diagnostic
}

func (b *Builder) Fatal(detail, summary string) *Diagnostics {
	return b.build(SeverityFatal, detail, summary)
}

func (b *Builder) Error(detail, summary string) *Diagnostics {
	return b.build(SeverityError, detail, summary)
}

func (b *Builder) Warn(detail, summary string) *Diagnostics {
	return b.build(SeverityWarn, detail, summary)
}

func (b *Builder) Info(detail, summary string) *Diagnostics {
	return b.build(SeverityInfo, detail, summary)
}

func (b *Builder) Debug(detail, summary string) *Diagnostics {
	return b.build(SeverityDebug, detail, summary)
}

func (b *Builder) Trace(detail, summary string) *Diagnostics {
	return b.build(SeverityTrace, detail, summary)
}

func (b *Builder) build(sev Severity, detail, summary string) *Diagnostics {
	diags, diag := b.d, b.e

	diag.Severity = sev
	diag.Detail = detail
	diag.Summary = summary

	diags.Append(*diag)

	return diags
}
