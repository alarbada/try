package try

import (
	"fmt"
	"runtime"
)

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
		if _, ok := arg.(error); ok {
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

func RecoverFunc(e *error, fn func(error)) {
	if recovered := recover(); recovered != nil {
		if err, ok := recovered.(error); ok {
			*e = err
		} else {
			*e = fmt.Errorf("unexpected panic: %v", recovered)
		}

		fn(*e)
	}
}

func panicWithFuncName(err error) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic(fmt.Errorf("unknown function: %w", err))
	}
	functionName := runtime.FuncForPC(pc).Name()
	panic(fmt.Errorf("%s: %w", functionName, err))
}
