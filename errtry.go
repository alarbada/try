package try

import (
	"errors"
	"fmt"
	"runtime"
)

type panicError struct {
	Function string
	Err      error
}

func (p panicError) Error() string {
	return fmt.Sprintf("%s: %v", p.Function, p.Err)
}

// Err will panic if err is not nil. To "catch" the panic, use `try.Recover`
func Err(err error) {
	if err != nil {
		panicWithFuncName(err)
	}
}

// Panicf forces a panic with the given msg and args.
func Panicf(msg string, args ...any) {
	panicWithFuncName(fmt.Errorf(msg, args...))
}

// Wrapf panics if an arg of type "error" is passed.
func Wrapf(msg string, args ...any) {
	for _, arg := range args {
		err, ok := arg.(error)
		if !ok {
			continue
		}

		if err != nil {
			panicWithFuncName(fmt.Errorf(msg, args...))
		}
	}
}

// Recover recovers from a panic and stores the error in e. If the recovered value is not an error, it will be wrapped in an error
func Recover(e *error) {
	if recovered := recover(); recovered != nil {
		if err, ok := recovered.(error); ok {
			*e = err
		} else {
			*e = fmt.Errorf("unexpected panic: %v", recovered)
		}
	}
}

func panicWithFuncName(err error) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic(panicError{"unknown function", err})
	}
	functionName := runtime.FuncForPC(pc).Name()
	panic(panicError{functionName, err})
}

func Unwrap(err error) error {
	var pErr panicError
	for errors.As(err, &pErr) {
		err = pErr.Err
	}
	return err
}
