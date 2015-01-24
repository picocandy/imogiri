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

	img, err := n.decode(r, opt.Format)
	if err != nil {
		return err
	}

	m := resize.Resize(opt.Width, opt.Height, img, resize.Bicubic)

	return n.encode(w, m, opt.Format)
}

func (n NFNT) decode(r io.Reader, format string) (m image.Image, err error) {
	switch format {
	case "jpg":
		m, err = jpeg.Decode(r)
	case "png":
		m, err = png.Decode(r)
	}

	return
}

func (n NFNT) encode(w io.Writer, m image.Image, format string) error {
	var err error

	switch format {
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
