package stacktrace

import (
	"errors"
	"strings"
)

type Error struct {
	err        error
	stackTrace StackTrace
}

func (te Error) Error() string {
	var builder strings.Builder
	builder.WriteString(te.err.Error())
	te.stackTrace.WriteOnStringBuilder(&builder, "\t")
	return builder.String()
}

func (te Error) Unwrap() error {
	return te.err
}

func WrapError(err error) *Error {
	var tracedError *Error
	if errors.As(err, &tracedError) {
		return tracedError
	}
	return &Error{
		err:        err,
		stackTrace: getStackTrace(3),
	}
}

func ExtractCause(err error) string {
	var tracedError *Error
	if errors.As(err, &tracedError) {
		return tracedError.err.Error()
	}
	return err.Error()
}
