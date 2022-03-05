package gerr

import (
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
		err:       err,
		grpcCodes: errType,
		stack:     populateStack(),
		query:     query,
		args:      args,
	}
}
