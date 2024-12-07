package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/alarbada/try"
)

func readJSONFile_WithoutTry(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", path, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON in file %s: %w", path, err)
	}

	return result, nil
}

func readJSONFile_WithTry(path string) (data map[string]interface{}, err error) {
	defer try.Recover(&err)

	file, err := os.Open(path)
	try.Wrapf("cannot open file %s: %w", path, err)
	defer file.Close()

	bytes, err := io.ReadAll(file)
	try.Err(err) // returning raw error from io.ReadAll

	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	try.Wrapf("invalid JSON in file %s: %w", path, err)

	return result, nil
}

func copyDir_WithoutTry(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, 0755); err != nil {
		return fmt.Errorf("failed creating destination dir %s: %w", dst, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		info, err := entry.Info()
		if err != nil {
			return err
		}

		if info.IsDir() {
			if err = copyDir_WithoutTry(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed copying directory %s: %w", srcPath, err)
			}
			continue
		}

		src, err := os.Open(srcPath)
		if err != nil {
			return fmt.Errorf("cannot open source file %s: %w", srcPath, err)
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return fmt.Errorf("cannot create destination file %s: %w", dstPath, err)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return nil
}

func copyDir_WithTry(src, dst string) (err error) {
	defer try.Recover(&err)

	entries, err := os.ReadDir(src)
	try.Err(err)

	err = os.MkdirAll(dst, 0755)
	try.Wrapf("failed creating destination dir %s: %w", dst, err)

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		info, err := entry.Info()
		try.Err(err)

		if info.IsDir() {
			err = copyDir_WithTry(srcPath, dstPath)
			try.Wrapf("failed copying directory %s: %w", srcPath, err)
			continue
		}

		src, err := os.Open(srcPath)
		try.Wrapf("cannot open source file %s: %w", srcPath, err)
		defer src.Close()

		dst, err := os.Create(dstPath)
		try.Wrapf("cannot create destination file %s: %w", dstPath, err)
		defer dst.Close()

		_, err = io.Copy(dst, src)
		try.Err(err)
	}

	return nil
}
