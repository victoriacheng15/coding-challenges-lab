package utils

import (
	"errors"
	"fmt"
	"os"
)

func ValidateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("no files or arguments provided")
	}

	return nil
}

func ReadContent(filename string, msg string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", msg, err)
		os.Exit(1)
	}

	return data
}
