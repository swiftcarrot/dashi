package model

import (
	"embed"
	"strings"

	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/flect/name"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

type presenter struct {
	Name        name.Ident
	Encoding    name.Ident
	Imports     []string
	Validations attrs.Attrs
}

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

	model := presenter{
		Name:        name.New(opts.Name.String()),
		Validations: validatable(opts.Attrs),
		Encoding:    name.New(opts.Encoding),
		Imports:     buildImports(opts),
	}

	ctx := map[string]interface{}{
		"opts":  opts,
		"model": model,
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
	g.Transformer(genny.Replace("$path$", opts.Path))
	g.Transformer(genny.Replace("$name$", opts.Name.Singularize().ToLower().String()))
	return g, nil
}
