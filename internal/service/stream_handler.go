package service

import (
	"context"
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

// handleStreamReponse handles the stream response, retries on errors, and updates the spinner
func handleStreamReponse[S StreamReceiverInterface[R], R any](stream S, requestCtx *context.Context, spinnerInstance *spinner.Spinner, reconnect ReconnectFunc[S], generateResponse GenerateResponse[R]) error {
	var message string
	var retries = 0
	errorChan := make(chan error)
	responseChan := make(chan R)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				response, err := stream.Recv()
				if err != nil {
					errorChan <- err
				}
				responseChan <- response
			}
		}
	}(*requestCtx)

	for {
		recvCtx, cancel := context.WithTimeout(*requestCtx, constant.Timeout)

		select {
		case <-recvCtx.Done():
			spinnerInstance.Stop()
			cancel()
			if !util.CanPerformRetry(retries, constant.MaxRetries) {
				return nil
			}
			var er error
			stream, er = reconnect()
			if er != nil {
				return nil
			}
			retries++
		case err := <-errorChan:
			spinnerInstance.Stop()
			cancel()
			if !util.IsRetryable(err) || !util.CanPerformRetry(retries, constant.MaxRetries) {
				if err != io.EOF {
					return err
				} else if err == io.EOF {
					log.Info(message)
				}
				return nil
			}
			stream, err = reconnect()
			if err != nil {
				return nil
			}
			retries++
		case response := <-responseChan:
			spinnerInstance.Stop()
			cancel()
			if !util.IsNil(response) {
				message = generateResponse(response)
				spinnerInstance.Prefix = fmt.Sprintf(" %s  ", message)
				spinnerInstance.Start()
				retries = 0
			}
		}
	}
}
