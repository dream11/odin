package constant

import "time"

// TraceID is the type for trace ID
type TraceID string

// VerboseEnabled is the type for verboseEnabledKey
type VerboseEnabled string

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

	// VerboseEnabledKey is the key used to store verbose value in context
	VerboseEnabledKey VerboseEnabled = "verbose"

	// VerboseFlag is the key used to store verbose value
	VerboseFlag string = "verbose"

	// LogLevelKey is the key used to set log level
	LogLevelKey = "ODIN_LOG_LEVEL"

	// ConsentMessageTemplate is the template for the consent message
	ConsentMessageTemplate = "\nYou are executing the above command on a restricted environment. Are you sure? Enter \033[1m%s\033[0m to continue:"

	// InitiatingRetryMessage is the message shown when retrying
	InitiatingRetryMessage = "Unable to reach Odin backend."

	// MaxRetries is the maximum number of retries
	MaxRetries = 5

	// Delay is the delay between retries
	Delay = 10 * time.Second

	// MaxRetriesReached is the message shown when max retries are reached
	MaxRetriesReached = "Max retries reached. Exiting...\nPlease check:\n- Your internet connection: \n- VPN connected properly"

	// RetryingMessage is the message shown when retrying
	RetryingMessage = "Retrying ... (%d/%d)"

	// CheckingAdditionalLogsMessage is the message shown when checking for additional logs
	CheckingAdditionalLogsMessage = "Execution completed. Checking for additional logs..."

	// ServiceExecutionMessageTemplate is the template for the service execution message
	ServiceExecutionMessageTemplate = "%s service: \u001B[41m %s \u001B[0m in environment: \u001B[41m %s \u001B[0m"

	// ComponentExecutionMessageTemplate is the template for the component execution message
	ComponentExecutionMessageTemplate = "%s component: \u001B[41m %s \u001B[0m in environment: \u001B[41m %s \u001B[0m"
)
