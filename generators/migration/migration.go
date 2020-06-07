package migration

import (
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("dashi:generators:migration", "../migration/templates")); err != nil {
		return g, err
	}

	g.Transformer(genny.Replace("-time-", opts.Time))
	g.Transformer(genny.Replace("-name-", opts.Name))

	return g, nil
}
