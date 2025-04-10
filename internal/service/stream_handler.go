package service

import (
	"fmt"
	"io"
	"time"

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
func handleStreamResponse[S StreamReceiverInterface[R], R any](stream S, reconnect ReconnectFunc[S], generateResponse GenerateResponse[R]) error {
	spinnerInstance := spinner.New(spinner.CharSets[constant.SpinnerType], constant.SpinnerDelay)
	err := spinnerInstance.Color(constant.SpinnerColor, constant.SpinnerStyle)
	if err != nil {
		return err
	}

	var message string
	for {
		response, err := stream.Recv()
		if err != nil {
			spinnerInstance.Stop()
			if err == io.EOF {
				log.Info(message)
				return nil
			}
			if !util.IsRetryable(err) {
				return err
			}
			stream, err = retryWithReconnect(reconnect, constant.MaxConnectRetries, constant.ConnectionRetryTimeout)
			if err != nil {
				return err
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

func retryWithReconnect[S any](reconnect ReconnectFunc[S], maxRetries int, retryTimeout time.Duration) (S, error) {
	var stream S
	var err error

	for retries := 0; retries < maxRetries; retries++ {

		if retries == 0 {
			log.Warnf(constant.InitiatingRetryMessage)
		}

		if retries < maxRetries {
			log.Infof(constant.RetryingMessage, retries+1, maxRetries)
		}

		stream, err = reconnect()
		if err == nil {
			return stream, nil
		}

		time.Sleep(retryTimeout)
	}
	log.Errorf("%s", constant.MaxRetriesReached)

	return stream, err
}
