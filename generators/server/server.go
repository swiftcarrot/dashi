package server

import (
	"embed"
	"text/template"

	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

type Options struct {
	Name    flect.Ident
	Package string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}

//go:embed templates/api templates/cmd
var templates embed.FS

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	err := g.Templates(&templates)
	if err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	return g, nil
}
