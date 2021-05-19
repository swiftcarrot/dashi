package model

import (
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/genny"
	"github.com/swiftcarrot/dashi/genny/gogen"
	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/flect/name"
)

func New(mops *Options) (*genny.Generator, error) {
	g := genny.New()
	if err := mops.Validate(); err != nil {
		return g, err
	}
	if err := g.Box(packr.New("dashi:generators:model", "../model/templates")); err != nil {
		return g, err
	}

	m := presenter{
		Name:        name.New(mops.Name.String()),
		Validations: validatable(mops.Attrs),
		Encoding:    name.New(mops.Encoding),
		Imports:     buildImports(mops),
	}

	ctx := map[string]interface{}{
		"opts":  mops,
		"model": m,
	}
	help := map[string]interface{}{
		"capitalize": flect.Capitalize,
		"trim_package": func(t string) string {
			i := strings.LastIndex(t, ".")
			if i == -1 {
				return t
			}
			return t[i+1:]
		},
	}

	t := gogen.TemplateTransformer(ctx, help)
	g.Transformer(t)
	g.Transformer(genny.Replace("-name-", mops.Name.Singularize().ToLower().String()))
	g.Transformer(genny.Replace("-path-", mops.Path))
	return g, nil
}
