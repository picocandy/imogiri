package imogiri

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNFNT_Resize_missingFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, Format: ""}
	n := NFNT{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := n.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "Please specify the format of the image", err.Error())
}

func TestNFNT_Resize_unknownFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, Format: "box"}
	n := NFNT{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := n.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, `Format "box" is not supported`, err.Error())
}

func TestNFNT_Resize_invalidImage(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, Format: "jpg"}
	n := NFNT{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := n.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, `unexpected EOF`, err.Error())
}

func TestNFNT_Resize(t *testing.T) {
	f := map[string]string{
		"png": "gopher.png",
		"jpg": "gopher.jpg",
	}

	for k, v := range f {
		x := ResizeOption{Width: 80, Height: 80, Format: k}
		n := NFNT{}
		s := new(bytes.Buffer)
		r := loadFixture(v)

		err := n.Resize(s, r, x)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, s.Len())
	}
}

func BenchmarkNFNT_Resize(b *testing.B) {
	n := NFNT{}
	r := loadFixture("gopher.jpg")
	x := ResizeOption{Width: 80, Height: 80, Format: "jpg"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := new(bytes.Buffer)
		r.Seek(0, 0)
		n.Resize(s, r, x)
	}
}
