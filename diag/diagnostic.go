package diag

import (
	"fmt"
	"strings"
)

// Diagnostic provides a singular point of diagnostic information to help
// diagnose program execution.
type Diagnostic struct {
	Severity  Severity
	Component Component
	Path      string
	Summary   string
	Detail    string
}

// Error implements error.
func (d Diagnostic) Error() string {
	var buf strings.Builder

	_, _ = fmt.Fprintf(&buf, "%s:", strings.ToUpper(d.Severity.String()))

	if component := d.Component.String(); componentInvalid.String() != component {
		_, _ = fmt.Fprintf(&buf, " (%s)", component)
	}

	if d.Path != "" {
		_, _ = fmt.Fprintf(&buf, " [%s]", d.Path)
	}

	_, _ = fmt.Fprintf(&buf, " Summary: \"%s\"", d.Summary)
	if d.Detail != "" {
		_, _ = fmt.Fprintf(&buf, " Detail: \"%s\"", d.Detail)
	}

	_, _ = fmt.Fprintln(&buf)

	return buf.String()
}
