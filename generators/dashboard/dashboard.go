package dashboard

import (
	"os/exec"
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/genny"
	"github.com/swiftcarrot/dashi/genny/gogen"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	if err := g.Box(packr.New("dashi:generators:dashboard", "../dashboard/templates")); err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	g.Transformer(genny.Dot())
	g.Command(exec.Command("yarn"))

	return g, nil
}
