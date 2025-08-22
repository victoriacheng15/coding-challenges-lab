package cmd

import (
	"testing"
)

func TestFormatLine(t *testing.T) {
	tests := []struct {
		name     string
		flags    Flags
		input    string
		lineNum  int
		expected string
	}{
		{
			name:     "No flags",
			flags:    Flags{},
			input:    "hello",
			lineNum:  1,
			expected: "hello",
		},
		{
			name:     "Number all lines (-n)",
			flags:    Flags{NumberLines: true},
			input:    "hello",
			lineNum:  1,
			expected: "     1\thello",
		},
		{
			name:     "Number non-blank (-b) non-blank line",
			flags:    Flags{NumberNonBlank: true},
			input:    "hello",
			lineNum:  1,
			expected: "     1\thello",
		},
		{
			name:     "Number non-blank (-b) blank line",
			flags:    Flags{NumberNonBlank: true},
			input:    "",
			lineNum:  1,
			expected: "",
		},
		{
			name:     "Show ends (-E)",
			flags:    Flags{ShowEnds: true},
			input:    "hello",
			lineNum:  1,
			expected: "hello$",
		},
		{
			name:     "Number all lines and show ends (-nE)",
			flags:    Flags{NumberLines: true, ShowEnds: true},
			input:    "hello",
			lineNum:  1,
			expected: "     1\thello$",
		},
		{
			name:     "Number non-blank and show ends (-bE) blank line",
			flags:    Flags{NumberNonBlank: true, ShowEnds: true},
			input:    "",
			lineNum:  1,
			expected: "$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			catFlags = tt.flags
			ln := tt.lineNum
			got := formatLine(tt.input, &ln)
			if got != tt.expected {
				t.Errorf("formatLine() = %q, want %q", got, tt.expected)
			}
		})
	}
}
