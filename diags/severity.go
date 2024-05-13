package diags

// Severity specifies the diagnostic logging level.
type Severity int

const (
	// severityInvalid represents an undefined severity and should not be used.
	severityInvalid Severity = iota
	// SeverityFatal provides information on a panic condition, where execution
	// must abort.
	SeverityFatal
	// SeverityError provides information on a state or condition that is
	// unexpected and/or needs to be rectified to resume normal execution.
	SeverityError
	// SeverityWarn provides information on a state or condition that may be
	// outside of normal execution parameters, but may be expected.
	SeverityWarn
	// SeverityInfo provides confirmation that the program is working as
	// expected.
	SeverityInfo
	// SeverityDebug provides diagnostic information of use when debugging a
	// program.
	SeverityDebug
	// SeverityTrace provides diagnostic information in an attempt to discover
	// all execution steps of a program.
	// This level is intended to be very verbose.
	SeverityTrace
)

// String implements the Stringer.
func (l Severity) String() string {
	switch l {
	case SeverityFatal:
		return "Fatal"
	case SeverityError:
		return "Error"
	case SeverityWarn:
		return "Warn"
	case SeverityInfo:
		return "Info"
	case SeverityDebug:
		return "Debug"
	case SeverityTrace:
		return "Trace"
	default:
		return "Invalid"
	}
}
