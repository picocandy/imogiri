package imogiri

import (
	"github.com/nfnt/resize"
	"github.com/picocandy/imogiri/quantizer"
	"golang.org/x/image/webp"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type NFNT struct{}

func (n *NFNT) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	err := n.formatChecker(opt.sourceFormat(), opt.Format)
	if err != nil {
		return err
	}

	if opt.sourceFormat() == "gif" && opt.Format == "gif" {
		img, err := gif.DecodeAll(r)
		if err != nil {
			return err
		}

		firstFrame := img.Image[0]
		im := image.NewRGBA(firstFrame.Bounds())

		for i, frame := range img.Image {
			bounds := frame.Bounds()
			draw.Draw(im, bounds, frame, bounds.Min, draw.Src)
			img.Image[i] = n.imageToPaletted(n.resizeImage(im, opt.Width, opt.Height, resize.Bilinear))
		}

		return gif.EncodeAll(w, img)
	}

	img, err := n.decode(r, opt.sourceFormat())
	if err != nil {
		return err
	}

	m := n.resizeImage(img, opt.Width, opt.Height, resize.Bicubic)

	return n.encode(w, m, opt.Format)
}

func (n *NFNT) decode(r io.Reader, format string) (m image.Image, err error) {
	switch format {
	case "jpg":
		m, err = jpeg.Decode(r)
	case "png":
		m, err = png.Decode(r)
	case "gif":
		m, err = gif.Decode(r)
	case "webp":
		m, err = webp.Decode(r)
	}

	return
}

func (n *NFNT) encode(w io.Writer, m image.Image, format string) error {
	var err error

	switch format {
	case "jpg":
		err = jpeg.Encode(w, m, &jpeg.Options{})
	case "png":
		err = png.Encode(w, m)
	case "gif":
		err = gif.Encode(w, m, &gif.Options{})
	}

	return err
}

func (n *NFNT) formatChecker(source, target string) error {
	return formatChecker(n.sourceFormats(), n.targetFormats(), source, target)
}

func (n *NFNT) sourceFormats() []string {
	return []string{"jpg", "png", "gif", "webp"}
}

func (n *NFNT) targetFormats() []string {
	return []string{"jpg", "png", "gif"}
}

func (n *NFNT) MatrixFormats() []string {
	return buildMatrix(n.sourceFormats(), n.targetFormats())
}

func (n *NFNT) SupportedActions() []Action {
	return []Action{ResizeAction}
}

func (n *NFNT) Name() string {
	return "NFNT"
}

func (n *NFNT) resizeImage(img image.Image, width, height uint, interpolation resize.InterpolationFunction) image.Image {
	return resize.Resize(width, height, img, interpolation)
}

func (n *NFNT) imageToPaletted(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)

	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, nil)
		q := &quantizer.MedianCutQuantizer{NumColor: 256}
		q.Quantize(pm, b, img, image.ZP)
	}

	return pm
}
