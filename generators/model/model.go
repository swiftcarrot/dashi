package model

import (
	"embed"
	"strings"

	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/flect/name"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

//go:embed templates
var templates embed.FS

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Templates(&templates); err != nil {
		return g, err
	}

	m := presenter{
		Name:        name.New(opts.Name.String()),
		Validations: validatable(opts.Attrs),
		Encoding:    name.New(opts.Encoding),
		Imports:     buildImports(opts),
	}

	ctx := map[string]interface{}{
		"opts":  opts,
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
	g.Transformer(genny.Replace("name", opts.Name.Singularize().ToLower().String()))
	g.Transformer(genny.Replace("path", opts.Path))
	return g, nil
}
