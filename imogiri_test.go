package imogiri

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type ResizerTest struct{}
type ResizerTestFailure struct{}

func (z *ResizerTest) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	return nil
}

func (z *ResizerTestFailure) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	return errors.New("Unable to resize the image")
}

func TestImogiri_Resize_noEngine(t *testing.T) {
	g := &Imogiri{ResizerEngines: []Resizer{}}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	x := ResizeOption{}

	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "No registered engine for resizing image", err.Error())
}

func TestImogiri_Resize_failure(t *testing.T) {
	g := &Imogiri{ResizerEngines: []Resizer{&ResizerTestFailure{}}}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	x := ResizeOption{}

	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "Unable to resize the image", err.Error())
}

func TestImogiri_Resize(t *testing.T) {
	g := &Imogiri{ResizerEngines: []Resizer{&ResizerTest{}}}
	s := new(bytes.Buffer)
	r := loadFixture("gopher.png")
	x := ResizeOption{Width: 80, Height: 80, Format: "png"}

	err := g.Resize(s, r, x)
	assert.Nil(t, err)
}

func loadFixture(name string) *bytes.Reader {
	fname := filepath.Join("fixtures", name)
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(fmt.Sprintf("Unable to open fixture file!. %s", fname))
	}

	return bytes.NewReader(b)
}
