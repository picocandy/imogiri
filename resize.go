package imogiri

import (
	"io"
	"strings"
)

type ResizeOption struct {
	Width      uint
	Height     uint
	FromFormat string
	Format     string
}

func (o ResizeOption) sourceFormat() string {
	if o.FromFormat == "" {
		return o.Format
	}

	return o.FromFormat
}

func (o ResizeOption) matrixFormat() string {
	return strings.Join([]string{o.sourceFormat(), o.Format}, ":")
}

type Resizer interface {
	Resize(w io.Writer, r io.Reader, opt ResizeOption) error
	MatrixFormats() []string
}
