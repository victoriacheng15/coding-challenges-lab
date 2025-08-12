package main

import (
	"testing"

	"wc-go/utils"
)

// Expected counts for sampleText:
// - bytes: 49
// - lines: 3 (three '\n')
// - words: 7
// - chars: 49 (ASCII only)

func TestBasicCounts(t *testing.T) {
	const sampleText = "Hello world!\nGo is fun.\nTest-driven development.\n"
	data := []byte(sampleText)

	tests := []struct {
		name    string
		countFn func([]byte) int
		want    int
	}{
		{"bytes", utils.CountBytes, 49},
		{"lines", utils.CountLines, 3},
		{"words", utils.CountWords, 7},
		{"chars", utils.CountChars, 49},
		{"ascii chars", utils.CountChars, 49},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.countFn(data)
			if got != tt.want {
				t.Fatalf("%s = %d; want %d", tt.name, got, tt.want)
			}
		})
	}

	t.Run("default format", func(t *testing.T) {
		want := "3 7 49 test.txt"
		got := utils.FormatDefault(data, "test.txt")
		if got != want {
			t.Fatalf("FormatDefault = %q; want %q", got, want)
		}
	})
}

// A Unicode sanity check so -m (chars) diverges from -c (bytes).
func TestUnicodeCounts(t *testing.T) {
	const u = "héllø 👋\n"
	data := []byte(u)
	tests := []struct {
		name    string
		countFn func([]byte) int
		want    int
	}{
		{"bytes", utils.CountBytes, 13},
		{"chars", utils.CountChars, 8},
		{"words", utils.CountWords, 2},
		{"lines", utils.CountLines, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.countFn(data)
			if got != tt.want {
				t.Fatalf("%s = %d; want %d", tt.name, got, tt.want)
			}
		})
	}
}

func TestStdinInput(t *testing.T) {
	const stdinSample = "line 1\nline 2\nline 3\n"
	data := []byte(stdinSample)

	tests := []struct {
		name    string
		countFn func([]byte) int
		want    int
	}{
		{"bytes", utils.CountBytes, 21},
		{"lines", utils.CountLines, 3},
		{"words", utils.CountWords, 6},
		{"chars", utils.CountChars, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.countFn(data)
			if got != tt.want {
				t.Fatalf("%s = %d; want %d", tt.name, got, tt.want)
			}
		})
	}

	t.Run("default format with stdin", func(t *testing.T) {
		want := "3 6 21 "
		got := utils.FormatDefault(data, "")
		if got != want {
			t.Fatalf("FormatDefault = %q; want %q", got, want)
		}
	})
}
