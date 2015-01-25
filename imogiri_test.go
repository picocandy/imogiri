package imogiri

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

type EngineTest struct{}
type EngineTestAlt struct{ EngineTest }
type EngineTestFailure struct{ EngineTest }

func (e *EngineTest) Name() string                                            { return "FirstEngine" }
func (e *EngineTest) MatrixFormats() []string                                 { return []string{"png:png"} }
func (e *EngineTest) SupportedActions() []Action                              { return []Action{ResizeAction} }
func (e *EngineTest) Resize(w io.Writer, r io.Reader, opt ResizeOption) error { return nil }

func (e *EngineTestAlt) Name() string { return "SecondEngine" }

func (e *EngineTestFailure) Name() string { return "EngineFailure" }
func (e *EngineTestFailure) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	return errors.New("Unable to resize the image")
}

func TestNewImogiri_noEngine(t *testing.T) {
	g := NewImogiri()

	assert.Equal(t, map[string]Engineer(nil), g.engines)
	assert.Equal(t, map[string][]string(nil), g.formatMatrix)
	assert.Equal(t, map[Action][]string(nil), g.actionMatrix)
}

func TestNewImogiri(t *testing.T) {
	n := &EngineTest{}
	a := &EngineTestAlt{}
	g := NewImogiri(n, a)

	assert.Equal(t, map[string]Engineer{"FirstEngine": n, "SecondEngine": a}, g.engines)
	assert.Equal(t, map[string][]string{"png:png": []string{"FirstEngine", "SecondEngine"}}, g.formatMatrix)
	assert.Equal(t, map[Action][]string{ResizeAction: []string{"FirstEngine", "SecondEngine"}}, g.actionMatrix)
}

func TestImogiri_RegisterEngine_multipleRegistration(t *testing.T) {
	n := &EngineTest{}
	g := &Imogiri{}

	err := g.RegisterEngine(n)
	assert.Nil(t, err)

	err = g.RegisterEngine(n)
	assert.NotNil(t, err)
	assert.Equal(t, `Engine "FirstEngine" already registered`, err.Error())
}

func TestImogiri_RegisterEngine(t *testing.T) {
	n := &EngineTest{}
	g := &Imogiri{}
	g.RegisterEngine(n)

	assert.Equal(t, map[string]Engineer{"FirstEngine": n}, g.engines)
	assert.Equal(t, map[string][]string{"png:png": []string{"FirstEngine"}}, g.formatMatrix)
	assert.Equal(t, map[Action][]string{ResizeAction: []string{"FirstEngine"}}, g.actionMatrix)
}

func TestImogiri_Resize_noEngine(t *testing.T) {
	g := &Imogiri{}
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	x := ResizeOption{}

	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "No registered engine for action: RESIZE", err.Error())
}

func TestImogiri_Resize_failure(t *testing.T) {
	n := &EngineTestFailure{}
	g := NewImogiri(n)
	s := new(bytes.Buffer)
	r := new(bytes.Reader)
	x := ResizeOption{Format: "png"}

	err := g.Resize(s, r, x)
	assert.NotNil(t, err)
	assert.Equal(t, "Unable to resize the image", err.Error())
}

func TestImogiri_Resize(t *testing.T) {
	n := &EngineTest{}
	g := NewImogiri(n)

	s := new(bytes.Buffer)
	r := loadFixture("gopher.png")
	x := ResizeOption{Width: 80, Height: 80, Format: "png"}

	err := g.Resize(s, r, x)
	assert.Nil(t, err)
}
