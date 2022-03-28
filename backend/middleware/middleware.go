package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc/metadata"

	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
)

type label string

const (
	endpointNameLabel label = "endpoint"
	requestLabel      label = "request"
	errLabel          label = "err"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				request := ctx.Value(requestLabel)
				if err != nil {
					_ = logger.Log(errLabel, gerr.LogErr(err), "took", time.Since(begin), "request", request)
				} else {
					_ = logger.Log("took", time.Since(begin), "request", request)
				}
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
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				var token string
				for _, i := range md.Get(string(jwt.JWTContextKey)) {
					token = i
				} // TODO: find out why can't just access using index
				ctx = context.WithValue(ctx, jwt.JWTContextKey, token)
			}
			return next(ctx, request)
		}
	}
}
