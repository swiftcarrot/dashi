package model

import (
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/flect/name"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
)

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()

	mops := &Options{
		Name:                   opts.Name,
		Attrs:                  opts.Attrs,
		Path:                   "graphql/model",
		Package:                "model",
		TestPackage:            "model",
		Encoding:               "json",
		ForceDefaultID:         true,
		ForceDefaultTimestamps: true,
	}
	if err := mops.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("scaffold:model:template", "../model/templates")); err != nil {
		return g, err
	}

	m := presenter{
		Name:        name.New(opts.Name.String()),
		Validations: validatable(opts.Attrs),
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
	g.Transformer(genny.Replace("-name-", mops.Name.Singularize().Dasherize().String()))
	g.Transformer(genny.Replace("-path-", mops.Path))
	return g, nil
}
