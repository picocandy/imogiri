package imogiri

import (
	"errors"
	"fmt"
)

func formatChecker(sourceFormats, targetFormats []string, source, target string) error {
	err := formatValidator(source, sourceFormats)
	if err != nil {
		return err
	}

	err = formatValidator(target, targetFormats)
	if err != nil {
		return err
	}

	return nil
}

func formatValidator(str string, formats []string) error {
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
