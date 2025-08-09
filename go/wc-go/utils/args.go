package utils

import "errors"

func ValidateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("no files or arguments provided")
	}

	return nil
}
