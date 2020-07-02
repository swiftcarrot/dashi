package webpacker

import (
	"os/exec"
	"text/template"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("webpacker/templates", "../webpacker/templates")); err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Replace("-name-", opts.Name.String()))
	g.Command(exec.Command("yarn"))

	return g, nil
}
