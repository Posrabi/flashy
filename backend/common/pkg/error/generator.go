package gerr

import (
	"google.golang.org/grpc/codes"
)

// Generic error wrapper that satisfies gRPC error codes.
func NewError(err error, errType codes.Code) error {
	return &Error{
		err:       err,
		grpcCodes: errType,
		stack:     populateStack(),
	}
}

// Error wrapper for scylla queries that satisfies gRPC error codes.
func NewScError(err error, errType codes.Code, query string, args []interface{}) error {
	return &Error{
		err:       err,
		grpcCodes: errType,
		stack:     populateStack(),
		query:     query,
		args:      args,
	}
}
