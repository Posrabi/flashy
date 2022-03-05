package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type label string

const (
	endpointNameLabel label = "endpoint"
	requestLabel      label = "request"
	errLabel          label = "error"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				request := ctx.Value(requestLabel)
				_ = logger.Log("error", err, "took", time.Since(begin), "request", request)
			}(time.Now())
			return next(ctx, request)
		}
	}
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
