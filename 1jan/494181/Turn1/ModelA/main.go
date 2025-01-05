package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileSystemError is a custom error type that contains detailed information about file system errors.
type FileSystemError struct {
	Err    error
	Path   string
	Action string
}

func (e *FileSystemError) Error() string {
	return fmt.Sprintf("%v: %v on %s", e.Err.Error(), e.Action, e.Path)
}

func newFileSystemError(err error, path string, action string) *FileSystemError {
	return &FileSystemError{Err: err, Path: path, Action: action}
}

// ReadFile reads a file and returns its content and any errors.
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, newFileSystemError(err, path, "reading")
	}
	return data, nil
}

// WriteFile writes content to a file and returns any errors.
func WriteFile(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return newFileSystemError(err, path, "writing")
	}
	return nil
}

// RemoveFile removes a file and returns any errors.
func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return newFileSystemError(err, path, "removing")
	}
	return nil
}

func main() {
	// Example file path
	filePath := filepath.Join("test", "test.txt")

	// Create and write to a file
	err := WriteFile(filePath, []byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully.")

	// Read from the file
	data, err := ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File read successfully:", string(data))

	// Remove the file
	err = RemoveFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File removed successfully.")
}
