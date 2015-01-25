package imogiri

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func loadFixture(name string) *bytes.Reader {
	fname := filepath.Join("fixtures", name)
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(fmt.Sprintf("Unable to open fixture file!. %s", fname))
	}

	return bytes.NewReader(b)
}
