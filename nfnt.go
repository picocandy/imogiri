package imogiri

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type NFNT struct{}

func (n NFNT) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	err := formatChecker(opt.Format, n.supportedFormats())
	if err != nil {
		return err
	}

	var img image.Image

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
