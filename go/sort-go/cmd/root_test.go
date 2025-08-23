package cmd

import (
	"os"
	"reflect"
	"sort"
	"testing"
)

var removeDuplicatesTests = []struct {
	name  string
	input []string
	want  []string
}{
	{
		name:  "empty slice",
		input: []string{},
		want:  nil, // removeDuplicates returns nil for empty input
	},
	{
		name:  "no duplicates",
		input: []string{"a", "b", "c"},
		want:  []string{"a", "b", "c"},
	},
	{
		name:  "with duplicates",
		input: []string{"b", "a", "c", "a", "b"},
		want:  []string{"a", "b", "c"},
	},
	{
		name:  "all duplicates",
		input: []string{"a", "a", "a"},
		want:  []string{"a"},
	},
	{
		name:  "single element",
		input: []string{"a"},
		want:  []string{"a"},
	},
}

var randomizeTests = []struct {
	name  string
	input []string
}{
	{
		name:  "empty slice",
		input: []string{},
	},
	{
		name:  "single element",
		input: []string{"a"},
	},
	{
		name:  "multiple elements",
		input: []string{"a", "b", "c", "d", "e"},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	tests := removeDuplicatesTests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			got := removeDuplicates(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomize(t *testing.T) {
	tests := randomizeTests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			got := randomize(inputCopy)

			// For empty or single element slices, result should be the same
			if len(tt.input) <= 1 {
				if !reflect.DeepEqual(got, tt.input) {
					t.Errorf("randomize() = %v, want %v", got, tt.input)
				}
				return
			}

			// For multiple elements, check that all original elements are present
			if len(got) != len(tt.input) {
				t.Errorf("randomize() length = %v, want %v", len(got), len(tt.input))
				return
			}

			// Sort both slices to compare content (ignoring order)
			sortedGot := make([]string, len(got))
			copy(sortedGot, got)
			sort.Strings(sortedGot)

			sortedInput := make([]string, len(tt.input))
			copy(sortedInput, tt.input)
			sort.Strings(sortedInput)

			if !reflect.DeepEqual(sortedGot, sortedInput) {
				t.Errorf("randomize() content = %v, want same content as %v", sortedGot, sortedInput)
			}
		})
	}
}

func TestProcessContent(t *testing.T) {
	// Create a temporary test file
	tempFile, err := os.CreateTemp("", "test*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test content to the file
	testContent := "line1\nline2\nline3\n"
	if _, err := tempFile.WriteString(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	tests := []struct {
		name    string
		args    []string
		want    []string
		wantErr bool
	}{
		{
			name:    "read from file",
			args:    []string{tempFile.Name()},
			want:    []string{"line1", "line2", "line3"},
			wantErr: false,
		},
		{
			name:    "file not found",
			args:    []string{"nonexistent.txt"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no args (stdin test would require mocking)",
			args:    []string{},
			want:    nil,
			wantErr: false, // This would normally read from stdin
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip stdin test as it requires input mocking
			if len(tt.args) == 0 {
				t.Skip("Skipping stdin test - requires input mocking")
				return
			}

			got, err := processContent(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("processContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
