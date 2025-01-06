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

// CreateDir creates a new directory and returns an error if there is an issue.
func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return newFileSystemError(err, path, "creating")
	}
	return nil
}

// ListFilesInDir lists the files in a directory and returns an error if there is an issue.
func ListFilesInDir(path string) ([]os.FileInfo, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, newFileSystemError(err, path, "reading")
	}
	return files, nil
}

// RemoveDir removes a directory and returns an error if there is an issue.
func RemoveDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return newFileSystemError(err, path, "removing")
	}
	return nil
}
func main() {
	// Example file path
	filePath := filepath.Join("test", "test.txt")
	dirPath := filepath.Join("test", "subdir")

	// Create a file
	err := WriteFile(filePath, []byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully.")

	// Create a directory
	err = CreateDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory created successfully.")

	// List files in a directory
	files, err := ListFilesInDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Files in directory:")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// Remove the file and directory
	err = RemoveFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = RemoveDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File removed successfully.")
	fmt.Println("Directory removed successfully.")
}
