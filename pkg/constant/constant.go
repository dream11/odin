package constant

import "time"

// TraceID is the type for trace ID
type TraceID string

const (
	// TEXT type output format
	TEXT = "text"
	// JSON type output format
	JSON = "json"
	// SpinnerColor Defines color of spinner
	SpinnerColor = "fgHiBlue"

	// SpinnerStyle Defines style of spinner
	SpinnerStyle = "bold"

	// SpinnerType Defines type of spinner
	SpinnerType = 14

	// SpinnerDelay Defines spinner delay
	SpinnerDelay = 100 * time.Millisecond

	// TraceIDKey is the key used to store traceID in context
	TraceIDKey TraceID = "trace-id"

	// LogLevelKey is the key used to set log level
	LogLevelKey = "ODIN_DEBUG"
)
