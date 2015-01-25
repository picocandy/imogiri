package imogiri

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

type Gifsicle struct{}

func (g *Gifsicle) Name() string {
	return g.execName()
}

func (g *Gifsicle) MatrixFormats() []string {
	return buildMatrix(g.sourceFormats(), g.targetFormats())
}

func (g *Gifsicle) SupportedActions() []Action {
	return []Action{ResizeAction}
}

func (g *Gifsicle) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
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

func (g *Gifsicle) execName() string {
	return "gifsicle"
}

func (g *Gifsicle) formatChecker(source, target string) error {
	return formatChecker(g.sourceFormats(), g.targetFormats(), source, target)
}

func (g *Gifsicle) sourceFormats() []string {
	return []string{"gif"}
}

func (g *Gifsicle) targetFormats() []string {
	return []string{"gif"}
}
