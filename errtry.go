package errtry

import (
	"fmt"
	"reflect"
	"runtime"
)

func panicWithFuncName(err error, f any) {
	functionValue := reflect.ValueOf(f)
	functionPtr := functionValue.Pointer()
	functionName := runtime.FuncForPC(functionPtr).Name()
	panic(fmt.Errorf("%s: %w", functionName, err))
}

func Recover(e *error) {
	if recovered := recover(); recovered != nil {
		if err, ok := recovered.(error); ok {
			*e = err
		} else {
			*e = fmt.Errorf("unexpected panic: %v", recovered)
		}
	}
}

func Wrapf(msg string, args ...any) {
	for _, arg := range args {
		if _, ok := arg.(error); ok {
			panic(fmt.Errorf(msg, args...))
		}
	}
}

func Try11[P any](f func(P) error, val P) {
	if err := f(val); err != nil {
		panicWithFuncName(err, f)
	}
}

func Try12[P, Q any](f func(P) (Q, error), val P) Q {
	res, err := f(val)
	if err != nil {
		panicWithFuncName(err, f)
	}

	return res
}

func Try21[P, Q any](f func(P, Q) error, val1 P, val2 Q) {
	if err := f(val1, val2); err != nil {
		panicWithFuncName(err, f)
	}
}

func Try22[P, Q, R any](f func(P, Q) (R, error), v1 P, v2 Q) R {
	res, err := f(v1, v2)
	if err != nil {
		panicWithFuncName(err, f)
	}

	return res
}

func Try23[P, Q, R any](f func(P, Q) (R, error), v1 P, v2 Q) R {
	res, err := f(v1, v2)
	if err != nil {
		panicWithFuncName(err, f)
	}

	return res
}

func Try31[P, Q, R any](f func(P, Q, R) error, v1 P, v2 Q, v3 R) {
	err := f(v1, v2, v3)
	if err != nil {
		panicWithFuncName(err, f)
	}
}

func Try32[P, Q, R, S any](f func(P, Q, R) (S, error), v1 P, v2 Q, v3 R) S {
	res, err := f(v1, v2, v3)
	if err != nil {
		panicWithFuncName(err, f)
	}

	return res
}

func Try33[P, Q, R, S, T any](f func(P, Q, R) (S, T, error), v1 P, v2 Q, v3 R) (S, T) {
	res1, res2, err := f(v1, v2, v3)
	if err != nil {
		panicWithFuncName(err, f)
	}

	return res1, res2
}
