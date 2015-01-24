package imogiri

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

var (
	gifSicleSourceFormats = []string{"gif"}
	gifSicleTargetFormats = []string{"gif"}
)

type GIFSicle struct{}

func (g GIFSicle) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	err := g.formatChecker(opt.sourceFormat(), opt.Format)
	if err != nil {
		return err
	}

	var stderr bytes.Buffer

	cmd := exec.Command(g.execName(), "--resize", fmt.Sprintf("%dx%d", opt.Width, opt.Height))
	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return errors.New(stderr.String())
	}

	return nil
}

func (g GIFSicle) execName() string {
	return "gifsicle"
}

func (g GIFSicle) formatChecker(source, target string) error {
	return formatChecker(gifSicleSourceFormats, gifSicleTargetFormats, source, target)
}
