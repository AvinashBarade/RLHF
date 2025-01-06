package main

import (
	"fmt"

	"os"

	"path/filepath"
)

// FileSystemError is a custom error type that contains detailed information about file system errors.

type FileSystemError struct {
	Err error

	Path string

	Action string
}

func (e *FileSystemError) Error() string {

	return fmt.Sprintf("%v: %v on %s", e.Err.Error(), e.Action, e.Path)

}

func newFileSystemError(err error, path string, action string) *FileSystemError {

	return &FileSystemError{Err: err, Path: path, Action: action}

}

// CreateDirectory creates a directory and returns any errors.
func CreateDirectory(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return newFileSystemError(err, path, "creating")
	}
	return nil
}

// ListDirectory lists the contents of a directory and returns any errors.
func ListDirectory(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, newFileSystemError(err, path, "listing")
	}
	var entries []string
	for _, file := range files {
		entries = append(entries, file.Name())
	}
	return entries, nil
}

// RemoveDirectory removes a directory and its contents recursively, and returns any errors.
func RemoveDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return newFileSystemError(err, path, "removing")
	}
	return nil
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
	// Example directory path
	dirPath := filepath.Join("test", "test-dir")

	// Create and write to a file within the directory
	filePath := filepath.Join(dirPath, "test.txt")
	err := WriteFile(filePath, []byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully.")

	// Create a directory
	err = CreateDirectory(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory created successfully.")

	// List the directory contents
	entries, err := ListDirectory(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory contents:", entries)

	// Remove the file
	err = RemoveFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File removed successfully.")

	// Remove the directory
	err = RemoveDirectory(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory removed successfully.")
}
