package imogiri

import (
	"errors"
	"fmt"
)

func formatChecker(str string, formats []string) error {
	if str == "" {
		return errors.New("Please specify the format of the image")
	}

	if !isSupported(str, formats) {
		return fmt.Errorf("Format %q is not supported", str)
	}

	return nil
}

func isSupported(str string, formats []string) bool {
	for _, f := range formats {
		if f == str {
			return true
		}
	}

	return false
}
