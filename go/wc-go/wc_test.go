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
		// {"words", utils.CountWords, 7},
		// {"chars", utils.CountChars, 49},
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

// func TestCountChars_ASCII(t *testing.T) {
// 	data := []byte(sampleText)
// 	want := 49
// 	got := utils.CountChars(data)
// 	if got != want {
// 		t.Fatalf("CountChars = %d; want %d", got, want)
// 	}
// }

// // A Unicode sanity check so -m (chars) diverges from -c (bytes).
// func TestUnicodeCounts(t *testing.T) {
// 	const u = "héllø 👋\n"
// 	data := []byte(u)

// 	// bytes: h(1)+é(2)+l(1)+l(1)+ø(2)+space(1)+👋(4)+\n(1) = 13
// 	if got, want := utils.CountBytes(data), 13; got != want {
// 		t.Fatalf("CountBytes(u) = %d; want %d", got, want)
// 	}

// 	// chars (runes): h é l l ø ␠ 👋 ␤ = 8
// 	if got, want := utils.CountChars(data), 8; got != want {
// 		t.Fatalf("CountChars(u) = %d; want %d", got, want)
// 	}

// 	// words: "héllø" and "👋" => 2
// 	if got, want := utils.CountWords(data), 2; got != want {
// 		t.Fatalf("CountWords(u) = %d; want %d", got, want)
// 	}

// 	// lines: one '\n'
// 	if got, want := utils.CountLines(data), 1; got != want {
// 		t.Fatalf("CountLines(u) = %d; want %d", got, want)
// 	}
// }

// func TestDefaultOutputFormat(t *testing.T) {
// 	data := []byte(sampleText)
// 	want := "3 7 49 test.txt"
// 	got := utils.FormatDefault(data, "test.txt")
// 	if got != want {
// 		t.Fatalf("FormatDefault = %q; want %q", got, want)
// 	}
// }
