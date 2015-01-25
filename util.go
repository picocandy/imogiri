package imogiri

import (
	"errors"
	"fmt"
	"github.com/rakyll/magicmime"
	"strings"
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

func mimeBuffer(b []byte) (string, error) {
	m, err := magicmime.New(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_ERROR)
	if err != nil {
		return "", err
	}
	defer m.Close()

	return m.TypeByBuffer(b)
}

func buildMatrix(sourceFormats, targetFormats []string) []string {
	matrix := []string{}

	for _, s := range sourceFormats {
		for _, t := range targetFormats {
			matrix = append(matrix, strings.Join([]string{s, t}, ":"))
		}
	}

	return matrix
}

func actionString(action Action) string {
	switch action {
	case ResizeAction:
		return "RESIZE"
	}

	return ""
}
