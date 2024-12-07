package try

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

var ErrForSure = errors.New("I did indeed, error")

func doesError() error {
	return ErrForSure
}

func TestExpectedUsage(t *testing.T) {
	err := func() (err error) {
		defer Recover(&err)

		Err(doesError())
		return nil
	}()

	is := is.New(t)

	is.True(errors.Is(err, ErrForSure))
}

func TestTry(t *testing.T) {
	is := is.New(t)

	t.Run("no error", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()
		Err(nil)
	})

	t.Run("with error", func(t *testing.T) {
		is := is.New(t)

		ErrTest := errors.New("test error")

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but got none")
			} else {
				err, ok := r.(error)
				is.True(ok)
				is.True(errors.Is(err, ErrTest))
			}
		}()
		Err(ErrTest)
	})
}

func TestWrapf(t *testing.T) {
	is := is.New(t)

	t.Run("no error in args", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()
		Wrapf("no error %s", "here")
	})

	t.Run("with error in args", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but got none")
			} else {
				err, ok := r.(error)
				is.True(ok)
				is.Equal(err.Error(), "error occurred: test error")
			}
		}()
		Wrapf("error occurred: %s", errors.New("test error"))
	})
}
