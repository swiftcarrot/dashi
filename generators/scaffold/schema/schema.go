package schema

import (
	"embed"
	"strings"
	"text/template"

	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

//go:embed templates
var templates embed.FS

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	helpers := template.FuncMap{
		"getFieldName": func(field string) string {
			return strings.Split(field, ":")[0]
		},
		"getFieldType": func(field string) string {
			return strings.Split(field, ":")[1]
		},
		"pascalize":  flect.Pascalize,
		"camelize":   flect.Camelize,
		"underscore": flect.Underscore,
	}
	// Change to camel
	data := map[string]interface{}{
		"opts": opts,
	}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Replace("$entity$", opts.Name.Singularize().ToLower().String()))
	g.Transformer(genny.Replace("$path$", "schema"))

	if err := g.Templates(&templates); err != nil {
		return g, err
	}

	return g, nil
}
