package imogiri

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type ResizerTest struct{}

func (z *ResizerTest) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	return nil
}

func TestImogiri_Resize(t *testing.T) {
	g := &Imogiri{ResizerEngines: []Resizer{&ResizerTest{}}}
	b := bytes.NewBuffer([]byte{})
	r := loadFixture("gopher.png")
	x := ResizeOption{Width: 80, Height: 80, Format: "png"}

	err := g.Resize(b, r, x)
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
