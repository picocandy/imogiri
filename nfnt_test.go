package imogiri

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"image/png"
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
	x := ResizeOption{Width: 80, Height: 80, FromFormat: "png", Format: "box"}
	n := NFNT{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := n.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, `Format "box" is not supported`, err.Error())
}

func TestNFNT_Resize_unknownFromFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, FromFormat: "box", Format: "png"}
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

func TestNFNT_Resize_conversion(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, FromFormat: "jpg", Format: "png"}
	n := NFNT{}
	s := new(bytes.Buffer)
	r := loadFixture("gopher.jpg")

	err := n.Resize(s, r, x)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, s.Len())

	mm, err := mimeBuffer(s.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, "image/png", mm)

	cfg, err := png.DecodeConfig(s)
	assert.Nil(t, err)
	assert.Equal(t, 80, cfg.Width)
	assert.Equal(t, 80, cfg.Height)
}

func TestNFNT_Resize(t *testing.T) {
	f := map[string]string{
		"png":  "gopher.png",
		"jpg":  "gopher.jpg",
		"webp": "gopher.webp",
	}

	c := []string{"png", "jpg"}

	for k, v := range f {
		for i := range c {
			x := ResizeOption{Width: 80, Height: 80, FromFormat: k, Format: c[i]}
			n := NFNT{}
			s := new(bytes.Buffer)
			r := loadFixture(v)

			err := n.Resize(s, r, x)
			assert.Nil(t, err)
			assert.NotEqual(t, 0, s.Len())

			mm, err := mimeBuffer(s.Bytes())
			assert.Nil(t, err)

			switch c[i] {
			case "png":
				assert.Equal(t, "image/png", mm)
			case "jpg":
				assert.Equal(t, "image/jpeg", mm)
			}
		}
	}
}

func BenchmarkNFNT_Resize_JPG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.jpg", "jpg", "")
}

func BenchmarkNFNT_Resize_JPG2PNG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.jpg", "png", "jpg")
}

func BenchmarkNFNT_Resize_PNG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.png", "png", "")
}

func BenchmarkNFNT_Resize_PNG2JPG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.png", "jpg", "png")
}

func BenchmarkNFNT_Resize_WEBP2PNG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.webp", "png", "webp")
}

func BenchmarkNFNT_Resize_WEBP2JPG(b *testing.B) {
	benchmarkResizeFormat(b, "gopher.webp", "jpg", "webp")
}

func benchmarkResizeFormat(b *testing.B, fixture, format, fromFormat string) {
	if fromFormat == "" {
		fromFormat = format
	}

	n := NFNT{}
	r := loadFixture(fixture)
	x := ResizeOption{Width: 80, Height: 80, FromFormat: fromFormat, Format: format}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := new(bytes.Buffer)
		r.Seek(0, 0)
		n.Resize(s, r, x)
	}
}
