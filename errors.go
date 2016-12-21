// Package errors provides utility routines for errors.
package errors

import (
	"database/sql/driver"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Grpc calls grpc.Errorf(code, "%v", err) with the passed error and code.
func Grpc(code codes.Code, err error) error {
	return grpc.Errorf(code, "%v", err)
}

// New is an error type that should be used analogously to the stdlib's
// errors.New function. It's advantageous to use because we can use const
// errors.
type New string

// Error implements the error interface.
func (n New) Error() string {
	return string(n)
}

// Value implements driver.Valuer. It just returns the value from Error.
func (n New) Value() (driver.Value, error) {
	return n.Error(), nil
}

// Error is analogous to fmt.Errorf.
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// InvalidArg returns the provided error wrapped with a gRPC InvalidArgument
// error code.
func InvalidArg(err error) error {
	return Grpc(codes.InvalidArgument, err)
}

// Internal returns the provided error wrapped with a gRPC Internal error code.
func Internal(err error) error {
	return Grpc(codes.Internal, err)
}

// Unauth returns the provided error wrapped with a gRPC Unauthorized error
// code.
func Unauth(err error) error {
	return Grpc(codes.Unauthenticated, err)
}

// Forbidden returns the provided error wrapped with a gRPC Permission Denied
// error code.
func Forbidden(err error) error {
	return Grpc(codes.PermissionDenied, err)
}

// NotFound returns the provided error wrapped with a gRPC Not Found error
// code.
func NotFound(err error) error {
	return Grpc(codes.NotFound, err)
}

// Unavailable returns the provided error wrapped with a gRPC Unavailable error
// code.
func Unavailable(err error) error {
	return Grpc(codes.Unavailable, err)
}

// FailedPrecondition returns the provided error wrapped with a gRPC Failed
// Precondition error code.
func FailedPrecondition(err error) error {
	return Grpc(codes.FailedPrecondition, err)
}

// Denied returns the provided error wrapped with a gRPC Permission Denied
// error code.
func Denied(err error) error {
	return Grpc(codes.PermissionDenied, err)
}
