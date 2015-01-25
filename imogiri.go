package imogiri

import (
	"fmt"
	"io"
)

type Action int

const (
	ResizeAction Action = iota
)

type Engineer interface {
	Name() string
	MatrixFormats() []string
	SupportedActions() []Action
}

type Imogiri struct {
	engines        map[string]Engineer
	formatMatrix   map[string][]string
	actionMatrix   map[Action][]string
	resizerEngines map[string][]Resizer
}

func NewImogiri(engines ...Engineer) *Imogiri {
	i := &Imogiri{}

	for _, e := range engines {
		i.RegisterEngine(e)
	}

	return i
}

func (g *Imogiri) RegisterEngine(engine Engineer) error {
	if len(g.engines) == 0 {
		g.engines = make(map[string]Engineer)
	}

	if _, ok := g.engines[engine.Name()]; ok {
		return fmt.Errorf("Engine %q already registered", engine.Name())
	}

	g.engines[engine.Name()] = engine
	g.calcFormatMatrix(engine)
	g.calcActionMatrix(engine)
	return nil
}

func (g *Imogiri) calcFormatMatrix(engine Engineer) {
	if len(g.formatMatrix) == 0 {
		g.formatMatrix = make(map[string][]string)
	}

	for _, m := range engine.MatrixFormats() {
		if _, ok := g.formatMatrix[m]; !ok {
			g.formatMatrix[m] = []string{}
		}

		g.formatMatrix[m] = append(g.formatMatrix[m], engine.Name())
	}
}

func (g *Imogiri) calcActionMatrix(engine Engineer) {
	if len(g.actionMatrix) == 0 {
		g.actionMatrix = make(map[Action][]string)
	}

	for _, a := range engine.SupportedActions() {
		if _, ok := g.actionMatrix[a]; !ok {
			g.actionMatrix[a] = []string{}
		}

		g.actionMatrix[a] = append(g.actionMatrix[a], engine.Name())
	}
}

func (g *Imogiri) enginePicker(matrixFormat string, action Action) (Engineer, error) {
	var ae []string
	var fe []string
	var ok bool

	if ae, ok = g.actionMatrix[action]; !ok {
		return nil, fmt.Errorf("No registered engine for action: %s", actionString(action))
	}

	if fe, ok = g.formatMatrix[matrixFormat]; !ok {
		return nil, fmt.Errorf("No registered engine for format: %s", matrixFormat)
	}

	for _, a := range ae {
		for _, f := range fe {
			if a == f {
				return g.engines[f], nil
			}
		}
	}

	return nil, fmt.Errorf("No registered engine for handling %s action with format %s", actionString(action), matrixFormat)
}

func (g *Imogiri) Resize(w io.Writer, r io.Reader, opt ResizeOption) error {
	e, err := g.enginePicker(opt.matrixFormat(), ResizeAction)
	if err != nil {
		return err
	}

	err = e.(Resizer).Resize(w, r, opt)
	if err != nil {
		return err
	}

	return nil
}
