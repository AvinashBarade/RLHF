package person

import (
	"fmt"
	"strings"
	"testing"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) ParseName(input string) error {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid name format: %s", input)
	}
	p.FirstName = parts[0]
	p.LastName = parts[1]
	return nil
}

func TestPerson_ParseName(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantErr   bool
		wantFirst string
		wantLast  string
	}{
		{
			name:      "ShouldParseCorrectName",
			input:     "Alice Johnson",
			wantErr:   false,
			wantFirst: "Alice",
			wantLast:  "Johnson",
		},
		{
			name:      "ShouldFailWithSinglePartName",
			input:     "Bob",
			wantErr:   true,
			wantFirst: "",
			wantLast:  "",
		},
		{
			name:      "ShouldTrimLeadingSpaces",
			input:     " Charlie Brown",
			wantErr:   false,
			wantFirst: "Charlie",
			wantLast:  "Brown",
		},
		{
			name:      "ShouldTrimTrailingSpaces",
			input:     "David Parker ",
			wantErr:   false,
			wantFirst: "David",
			wantLast:  "Parker",
		},
		{
			name:      "ShouldHandleMultipleSpacesBetweenNames",
			input:     "Emma Rose Wood",
			wantErr:   false,
			wantFirst: "Emma",
			wantLast:  "Rose Wood",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Person{}
			err := p.ParseName(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseName() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if p.FirstName != tc.wantFirst {
				t.Errorf("ParseName() got FirstName = %v, want %v", p.FirstName, tc.wantFirst)
			}
			if p.LastName != tc.wantLast {
				t.Errorf("ParseName() got LastName = %v, want %v", p.LastName, tc.wantLast)
			}
		})
	}
}
