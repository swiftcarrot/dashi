package packages

import (
	"embed"
	"html/template"
	"os/exec"

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
