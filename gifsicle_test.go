package imogiri

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestGIFSicle_executable(t *testing.T) {
	g := GIFSicle{}
	path, err := exec.LookPath(g.execName())
	assert.Nil(t, err)
	assert.NotEmpty(t, path)
}

func TestGIFSicle_Resize_missingFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, Format: ""}
	g := GIFSicle{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "Please specify the format of the image", err.Error())
}

func TestGIFSicle_Resize_unknownFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, FromFormat: "gif", Format: "png"}
	g := GIFSicle{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, `Format "png" is not supported`, err.Error())
}

func TestGIFSicle_Resize_unknownFromFormat(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, FromFormat: "png", Format: "gif"}
	g := GIFSicle{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, `Format "png" is not supported`, err.Error())
}

func TestGIFSicle_Resize_invalidImage(t *testing.T) {
	x := ResizeOption{Width: 80, Height: 80, Format: "gif"}
	g := GIFSicle{}
	s := new(bytes.Buffer)
	r := loadFixture("gopher.jpg")
	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `file not in GIF format`)
}

func TestGIFSicle_Resize(t *testing.T) {
	f := []string{"gopher.gif", "animation.gif"}

	for i := range f {
		x := ResizeOption{Width: 80, Height: 80, Format: "gif"}
		g := GIFSicle{}
		s := new(bytes.Buffer)
		r := loadFixture(f[i])

		err := g.Resize(s, r, x)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, s.Len())

		mm, err := mimeBuffer(s.Bytes())
		assert.Nil(t, err)
		assert.Equal(t, "image/gif", mm)
	}
}

func BenchmarkGIFSicle_singleFrame(b *testing.B) {
	g := GIFSicle{}
	r := loadFixture("gopher.gif")
	x := ResizeOption{Width: 80, Height: 80, Format: "gif"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := new(bytes.Buffer)
		r.Seek(0, 0)
		g.Resize(s, r, x)
	}
}

func BenchmarkGIFSicle_multipleFrames(b *testing.B) {
	g := GIFSicle{}
	r := loadFixture("animation.gif")
	x := ResizeOption{Width: 80, Height: 80, Format: "gif"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := new(bytes.Buffer)
		r.Seek(0, 0)
		g.Resize(s, r, x)
	}
}
