package error

import "google.golang.org/grpc/codes"

type GRPCErrorType int

const (
	Default GRPCErrorType = iota + 1
	AuthFailed
	ArgsNotAccepted
	RequestFailed
	DataProcessError
	PermissionDenied
	Unimplemented
	NotFound
	Unavailable
	AlreadyExists
	Aborted
	ResourceExhausted
	RedisUnknownError
	KubernetesUnknownError
)

type GRPCError struct {
	StackTrace []uintptr
	goErr      error
	Query      string
	Args       []interface{}
	Type       GRPCErrorType
}

func (e *GRPCError) Error() string {
	return e.goErr.Error()
}

// GetGRPCCode converts our error type to corresponding GRPC error type.
func (e *GRPCError) GetGRPCCode() codes.Code { //nolint:gocyclo
	switch e.Type {
	case Default:
		return codes.Unknown
	case AuthFailed:
		return codes.Unauthenticated
	case NotFound:
		return codes.NotFound
	case ArgsNotAccepted:
		return codes.InvalidArgument
	case DataProcessError:
		return codes.Internal
	case RequestFailed:
		return codes.Unavailable
	case PermissionDenied:
		return codes.PermissionDenied
	case Unimplemented:
		return codes.Unimplemented
	case Unavailable:
		return codes.Unavailable
	case RedisUnknownError:
		return codes.Internal
	case AlreadyExists:
		return codes.AlreadyExists
	case Aborted:
		return codes.Aborted
	case ResourceExhausted:
		return codes.ResourceExhausted
	case KubernetesUnknownError:
		return codes.Internal
	}

	return codes.Unknown
}

func GetGRPCCode(err error) codes.Code {
	if grpcErr, ok := err.(*GRPCError); ok { // nolint: errorlint
		return grpcErr.GetGRPCCode()
	}
	return codes.Internal
}

// TODO: a lot more error handling and this is not right
