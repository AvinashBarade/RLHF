package main

import (
	"fmt"
	"os"
)

// UserProfile represents a user profile
type UserProfile struct {
	Name   string
	Age    int
	Skills []string
	Social map[string]string
}

func readFile(filePath string) ([]byte, error) {
	// ... (Same code as before)
}

func formatUserProfile(profile UserProfile) string {
	return fmt.Sprintf("User Profile: \n Name: %s\n Age: %d\n Skills: %v\n Social Media:\n  %v\n",
		profile.Name, profile.Age, profile.Skills, profile.Social)
}

func main() {
	filePath := "example.txt"
	data, err := readFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Example usage of UserProfile struct
	userProfile := UserProfile{
		Name:   "Alice",
		Age:    25,
		Skills: []string{"Go", "Python", "JavaScript"},
		Social: map[string]string{
			"GitHub":  "alice123",
			"Twitter": "@alice_dev",
		},
	}

	fmt.Println(formatUserProfile(userProfile))
}
