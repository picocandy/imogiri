package imogiri

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNFNT_Resize(t *testing.T) {
	f := map[string]string{
		"png": "gopher.png",
		"jpg": "gopher.jpg",
	}

	for k, v := range f {
		x := ResizeOption{Width: 80, Height: 80, Format: k}
		n := NFNT{}
		s := bytes.NewBuffer([]byte{})
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
		s := bytes.NewBuffer([]byte{})
		r.Seek(0, 0)
		n.Resize(s, r, x)
	}
}
