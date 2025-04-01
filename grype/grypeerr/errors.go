package grypeerr

var (
	// ErrAboveSeverityThreshold indicates when a vulnerability severity is discovered that is equal
	// or above the given --fail-on severity value.
	ErrAboveSeverityThreshold = NewExpectedErr("discovered vulnerabilities at or above the severity threshold")
)
