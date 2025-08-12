package utils

import (
	"fmt"
	"unicode"
)

type Flags struct {
	Lines bool
	Words bool
	Bytes bool
	Chars bool
}

var (
	WcFlags Flags
)

func CountBytes(data []byte) int {
	return len(data)
}

func CountLines(data []byte) int {
	lines := 0
	for _, b := range data {
		if b == '\n' {
			lines++
		}
	}
	if len(data) > 0 && data[len(data)-1] != '\n' {
		lines++
	}

	return lines
}

func CountWords(data []byte) int {
	inWord := false
	count := 0

	for _, item := range string(data) {
		if unicode.IsSpace(item) {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}

	return count
}

func CountChars(data []byte) int {
	return len([]rune(string(data)))
}

func FormatDefault(data []byte, name string) string {
	return fmt.Sprintf("%d %d %d %s",
		CountLines(data),
		CountWords(data),
		CountBytes(data),
		name)
}

func PrintCounts(data []byte, name string) {
	switch {
	case WcFlags.Bytes:
		fmt.Printf("%8d %s\n", CountBytes(data), name)
	case WcFlags.Lines:
		fmt.Printf("%8d %s\n", CountLines(data), name)
	case WcFlags.Words:
		fmt.Printf("%8d %s\n", CountWords(data), name)
	case WcFlags.Chars:
		fmt.Printf("%8d %s\n", CountChars(data), name)
	default:
		fmt.Println(FormatDefault(data, name))
	}
}
