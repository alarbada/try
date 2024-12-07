package try

import (
	"fmt"
	"runtime"
)

func Err(err error) {
	if err != nil {
		panicWithFuncName(err)
	}
}

func Wrapf(msg string, args ...any) {
	for _, arg := range args {
		if _, ok := arg.(error); ok {
			panic(fmt.Errorf(msg, args...))
		}
	}
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

func panicWithFuncName(err error) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic(fmt.Errorf("unknown function: %w", err))
	}
	functionName := runtime.FuncForPC(pc).Name()
	panic(fmt.Errorf("%s: %w", functionName, err))
}

type tryer struct{}

var Tryer tryer

func (tryer) Try(err error)                 { Err(err) }
func (tryer) Wrapf(msg string, args ...any) { Wrapf(msg, args...) }
func (tryer) Recover(e *error)              { Recover(e) }
