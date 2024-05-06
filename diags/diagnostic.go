package diags

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

	_, _ = fmt.Fprintf(&buf, "%s:", d.Severity)

	if component := d.Component.String(); componentInvalid.String() != component {
		_, _ = fmt.Fprintf(&buf, " (%s)", component)
	}

	if d.Path != "" {
		_, _ = fmt.Fprintf(&buf, " [%s]", d.Path)
	}

	_, _ = fmt.Fprintf(&buf, " %s", d.Summary)
	if d.Detail != "" {
		_, _ = fmt.Fprintf(&buf, "\n\n%s", d.Detail)
	}

	return buf.String()
}
