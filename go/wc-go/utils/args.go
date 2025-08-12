package utils

import (
	"errors"
	"fmt"
	"io"
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
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	return data
}

func ReadStdin() []byte {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	return data
}
