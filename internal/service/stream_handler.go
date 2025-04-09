package service

import (
	"fmt"
	"io"

	"github.com/briandowns/spinner"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/util"
	log "github.com/sirupsen/logrus"
)

// StreamReceiverInterface defines the Recv method that the stream must implement
type StreamReceiverInterface[R any] interface {
	Recv() (R, error)
}

// ReconnectFunc defines a function type for reconnecting the stream
type ReconnectFunc[S any] func() (S, error)

// GenerateResponse defines a function type for generating a response message
type GenerateResponse[R any] func(response R) string

// handleStreamResponse handles the stream response, retries on errors, and updates the spinner
func handleStreamResponse[S StreamReceiverInterface[R], R any](stream S, spinnerInstance *spinner.Spinner, reconnect ReconnectFunc[S], generateResponse GenerateResponse[R]) error {
	var message string
	var retries = 0
	for {
		response, err := stream.Recv()
		if err != nil {
			spinnerInstance.Stop()
			if err == io.EOF {
				log.Info(message)
				return nil
			}
			if !util.IsRetryable(err) || !util.CanPerformRetry(retries, constant.MaxRetries) {
				return err
			}
			retries++
			stream, err = reconnect()
			if err != nil {
				return nil
			}
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		} else {
			spinnerInstance.Stop()
			message = generateResponse(response)
			spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
			spinnerInstance.Start()
		}
	}
}
