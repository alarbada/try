package try

import (
	"errors"
	"fmt"
	"testing"
)

// Traditional error handling
func normalErrorHandling() (err error) {
	if err := stepOne(); err != nil {
		return err
	}
	if err := stepTwo(); err != nil {
		return err
	}
	if err := stepThree(); err != nil {
		return err
	}
	return nil
}

// Try package error handling
func tryErrorHandling() (err error) {
	defer Recover(&err)

	Err(stepOne())
	Err(stepTwo())
	Err(stepThree())

	return nil
}

// Mock functions that may return errors
func stepOne() error {
	return nil
}

func stepTwo() error {
	return nil
}

var ErrHappened = errors.New("an error happened")

var stepThree = func() error {
	return fmt.Errorf("%s: %w", "some error", ErrHappened)
}

// Benchmarks
func BenchmarkNormalErrorHandling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normalErrorHandling()
	}
}

func BenchmarkTryErrorHandling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = tryErrorHandling()
	}
}

// Benchmarks with no errors
func BenchmarkNormalErrorHandlingNoError(b *testing.B) {
	// Temporarily make stepThree return nil
	old := stepThree
	stepThree = func() error { return nil }
	defer func() { stepThree = old }()

	for i := 0; i < b.N; i++ {
		_ = normalErrorHandling()
	}
}

func BenchmarkTryErrorHandlingNoError(b *testing.B) {
	// Temporarily make stepThree return nil
	old := stepThree
	stepThree = func() error { return nil }
	defer func() { stepThree = old }()

	for i := 0; i < b.N; i++ {
		_ = tryErrorHandling()
	}
}
