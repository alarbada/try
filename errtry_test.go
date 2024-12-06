package errtry

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var errTest = errors.New("test error")

// Test helper functions
func successFunc1(s string) error {
	return nil
}

func errorFunc1(s string) error {
	return errTest
}

func successFunc2(s string) (int, error) {
	return 42, nil
}

func errorFunc2(s string) (int, error) {
	return 0, errTest
}

func successFunc3(a, b string) error {
	return nil
}

func errorFunc3(a, b string) error {
	return errTest
}

func successFunc4(a, b string) (int, error) {
	return 42, nil
}

func errorFunc4(a, b string) (int, error) {
	return 0, errTest
}

func successFunc5(a, b, c string) error {
	return nil
}

func errorFunc5(a, b, c string) error {
	return errTest
}

func successFunc6(a, b, c string) (int, error) {
	return 42, nil
}

func errorFunc6(a, b, c string) (int, error) {
	return 0, errTest
}

func successFunc7(a, b, c string) (int, string, error) {
	return 42, "success", nil
}

func errorFunc7(a, b, c string) (int, string, error) {
	return 0, "", errTest
}

// Tests for Try11
func TestTry11(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	func() {
		defer Recover(&err)
		Try11(successFunc1, "test")
	}()
	r.NoError(err)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try11(errorFunc1, "test")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc1")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try12
func TestTry12(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	var result int
	func() {
		defer Recover(&err)
		result = Try12(successFunc2, "test")
	}()
	r.NoError(err)
	r.Equal(42, result)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try12(errorFunc2, "test")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc2")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try21
func TestTry21(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	func() {
		defer Recover(&err)
		Try21(successFunc3, "test1", "test2")
	}()
	r.NoError(err)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try21(errorFunc3, "test1", "test2")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc3")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try22
func TestTry22(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	var result int
	func() {
		defer Recover(&err)
		result = Try22(successFunc4, "test1", "test2")
	}()
	r.NoError(err)
	r.Equal(42, result)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try22(errorFunc4, "test1", "test2")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc4")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try31
func TestTry31(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	func() {
		defer Recover(&err)
		Try31(successFunc5, "test1", "test2", "test3")
	}()
	r.NoError(err)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try31(errorFunc5, "test1", "test2", "test3")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc5")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try32
func TestTry32(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	var result int
	func() {
		defer Recover(&err)
		result = Try32(successFunc6, "test1", "test2", "test3")
	}()
	r.NoError(err)
	r.Equal(42, result)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try32(errorFunc6, "test1", "test2", "test3")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc6")
	r.Contains(err.Error(), errTest.Error())
}

// Tests for Try33
func TestTry33(t *testing.T) {
	r := require.New(t)

	// Test success case
	var err error
	var result1 int
	var result2 string
	func() {
		defer Recover(&err)
		result1, result2 = Try33(successFunc7, "test1", "test2", "test3")
	}()
	r.NoError(err)
	r.Equal(42, result1)
	r.Equal("success", result2)

	// Test error case
	err = nil
	func() {
		defer Recover(&err)
		Try33(errorFunc7, "test1", "test2", "test3")
	}()
	r.Error(err)
	r.Contains(err.Error(), "errorFunc7")
	r.Contains(err.Error(), errTest.Error())
}

// Test Recover itself
func TestRecover(t *testing.T) {
	r := require.New(t)

	// Test with error panic
	var err error
	func() {
		defer Recover(&err)
		panic(errors.New("test panic error"))
	}()
	r.Error(err)
	r.Contains(err.Error(), "test panic error")

	// Test with non-error panic
	err = nil
	func() {
		defer Recover(&err)
		panic("string panic")
	}()
	r.Error(err)
	r.Contains(err.Error(), "unexpected panic: string panic")
}
