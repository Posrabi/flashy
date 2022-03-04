package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				_ = logger.Log("error", err, "took", time.Since(begin))
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
