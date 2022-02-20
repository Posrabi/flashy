package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type label string

const (
	endpointNameLabel label = "endpoint"
	requestLabel      label = "request"
)

type ErrorLogHandler struct {
	logger log.Logger
}

func AddRequestToContext(name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			ctx = context.WithValue(ctx, endpointNameLabel, name)
			ctx = context.WithValue(ctx, requestLabel, request)
			return next(ctx, request)
		}
	}
}

func NewErrorLogHandler(logger log.Logger) *ErrorLogHandler {
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
