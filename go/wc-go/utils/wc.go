package utils

import "unicode"

func CountBytes(data []byte) int {
	return len(data)
}

func CountLines(data []byte) int {
	lines := 0
	for _, b := range data {
		if b == '\n' {
			lines++
		}

		if len(data) > 0 && data[len(data)-1] != '\n' {
			lines++
		}
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
	// TODO: implement
	return ""
}
