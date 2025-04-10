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
	LogLevelKey = "ODIN_LOG_LEVEL"

	// ConsentMessageTemplate is the template for the consent message
	ConsentMessageTemplate = "\nYou are executing the above command on a restricted environment. Are you sure? Enter \033[1m%s\033[0m to continue:"
	// InitiatingRetryMessage is the message shown when retrying
	InitiatingRetryMessage = "Unable to reach Odin backend."
	//FailedRetryMessage is the message shown when  client fails to connect to backend after repeated attempts
	FailedRetryMessage = "failed to connect to odin backend"
	// MaxConnectRetries is the maximum number of connection retries
	MaxConnectRetries = 10
	// ConnectionRetryTimeout is the timeout for connection retries
	ConnectionRetryTimeout = 5 * time.Second
	MaxRetriesReached      = "Max retries reached. Exiting...\nPlease check:\n- Your internet connection: \n- VPN connected properly"
	DescribeEnv            = "\n- Please use describe env command to check the status.\n \u001B[1m odin describe env --name %s\u001B[0m  "
	RetryingMessage        = "Retrying ... (%d/%d)"
)
