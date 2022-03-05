package gerr

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

func NewError(err error, errType codes.Code) error {
	return &Error{
		err:       err,
		grpcCodes: errType,
		stack:     populateStack(),
	}
}

func NewScError(err error, errType codes.Code, query string, args []interface{}) error {
	return &Error{
		err:       fmt.Errorf("query %s failed with args %v, error: %w", query, args, err),
		grpcCodes: errType,
		stack:     populateStack(),
	}
}
