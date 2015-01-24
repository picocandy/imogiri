package imogiri

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type NFNT struct{}

func (n NFNT) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	if opt.Format == "" {
		return errors.New("Please specify the format of the image")
	}

	if !n.isSupported(opt.Format) {
		return fmt.Errorf("Format %q is not supported", opt.Format)
	}

	var img image.Image
	var err error

	switch opt.Format {
	case "jpg":
		img, err = jpeg.Decode(r)
	case "png":
		img, err = png.Decode(r)
	}

	if err != nil {
		return err
	}

	m := resize.Resize(opt.Width, opt.Height, img, resize.Bicubic)

	switch opt.Format {
	case "jpg":
		err = jpeg.Encode(w, m, &jpeg.Options{})
	case "png":
		err = png.Encode(w, m)
	}

	return err
}

func (n NFNT) supportedFormats() []string {
	return []string{"png", "jpg"}
}

func (n NFNT) isSupported(format string) bool {
	for _, f := range n.supportedFormats() {
		if f == format {
			return true
		}
	}

	return false
}
