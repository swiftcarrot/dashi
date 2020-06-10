package model

import (
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("model/templates", "../model/templates")); err != nil {
		return g, err
	}
	return g, nil
}
