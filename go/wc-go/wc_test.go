package main

import (
	"testing"

	"wc-go/utils"
)

const sampleText = "Hello world!\nGo is fun.\nTest-driven development.\n"

// Expected counts for sampleText:
// - bytes: 49
// - lines: 3 (three '\n')
// - words: 7
// - chars: 49 (ASCII only)

func TestBasicCounts(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.countFn(data)
			if got != tt.want {
				t.Fatalf("%s = %d; want %d", tt.name, got, tt.want)
			}
		})
	}
	t.Run("ascii chars", func(t *testing.T) {
		got := utils.CountChars(data)
		want := 49
		if got != want {
			t.Fatalf("CountChars = %d; want %d", got, want)
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

// func TestDefaultOutputFormat(t *testing.T) {
// 	data := []byte(sampleText)
// 	want := "3 7 49 test.txt"
// 	got := utils.FormatDefault(data, "test.txt")
// 	if got != want {
// 		t.Fatalf("FormatDefault = %q; want %q", got, want)
// 	}
// }
