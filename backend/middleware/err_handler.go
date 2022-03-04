package middleware

import (
	"context"

	"github.com/go-kit/log"
)

type label string

const (
	endpointNameLabel label = "endpoint"
	requestLabel      label = "request"
)

type ErrorLogHandler struct {
	logger log.Logger
}

func NewLogErrorHandler(logger log.Logger) *ErrorLogHandler {
	return &ErrorLogHandler{
		logger: logger,
	}
}

func (e *ErrorLogHandler) Handle(ctx context.Context, err error) {
	if err == nil {
		return
	}

	var returnEndpoint string
	val := ctx.Value(endpointNameLabel)
	if val == nil {
		returnEndpoint = ""
	} else {
		returnEndpoint = val.(string)
	}

	request := ctx.Value(requestLabel)

	_ = e.logger.Log("err", err, "endpoint", returnEndpoint, "request", request)
}
