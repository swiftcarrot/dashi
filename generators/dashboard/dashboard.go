package dashboard

import (
	"embed"
	"os/exec"
	"text/template"

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

	g.Transformer(genny.Replace("$dot$", "."))
	// TODO: remove $layout$ replace
	g.Transformer(genny.Replace("$layout$", "__layout__"))
	g.Command(exec.Command("yarn"))

	return g, nil
}
