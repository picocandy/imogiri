package imogiri

import (
	"errors"
	"io"
	"math/rand"
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

type Resizer interface {
	Resize(w io.Writer, r io.Reader, opt ResizeOption) error
}

type Imogiri struct {
	ResizerEngines []Resizer
}

func (i *Imogiri) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	e, err := i.resizer()
	if err != nil {
		return err
	}

	err = e.Resize(w, r, opt)
	if err != nil {
		return err
	}

	return nil
}

func (i *Imogiri) resizer() (Resizer, error) {
	n := len(i.ResizerEngines)

	if n == 0 {
		return nil, errors.New("No registered engine for resizing image")
	}

	return i.ResizerEngines[rand.Intn(n)], nil
}
