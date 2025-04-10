package retry

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// RetryConfig holds the configuration for retry behavior
type RetryConfig struct {
	// MaxRetries is the maximum number of retry attempts
	MaxRetries int
	// RetryDelay is the delay between retry attempts
	RetryDelay time.Duration
	// RetryableError is a function that determines if an error is retryable
	RetryableError func(error) bool
}

// DefaultRetryConfig returns a default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:     3,
		RetryDelay:     time.Second,
		RetryableError: func(err error) bool { return true },
	}
}

// Retry executes the given function and retries if it fails based on the provided configuration
func Retry[T any](fn func() (T, error), config *RetryConfig) (T, error) {
	var result T
	var err error

	if config == nil {
		config = DefaultRetryConfig()
	}

	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		if attempt > 0 {
			log.Infof("Retry attempt %d/%d", attempt, config.MaxRetries)
			time.Sleep(config.RetryDelay)
		}

		result, err = fn()
		if err == nil {
			return result, nil
		}

		if !config.RetryableError(err) {
			return result, err
		}

		if attempt == config.MaxRetries {
			log.Errorf("Max retries (%d) reached. Last error: %v", config.MaxRetries, err)
			return result, err
		}
	}

	return result, err
}
