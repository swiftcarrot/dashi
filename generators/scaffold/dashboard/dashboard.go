package dashboard

import (
	"embed"
	"os/exec"
	"text/template"

	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

//go:embed templates/$pages$ templates/$components$
var templates embed.FS

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	name := opts.Name.Pascalize().Pluralize().ToLower().String()

	g.Transformer(genny.Replace("$pages$", "packages/dashboard/src/pages/"+name))
	g.Transformer(genny.Replace("$components$", "packages/components/src/"+name))

	if err := g.Templates(&templates); err != nil {
		return g, err
	}

	g.Command(exec.Command("yarn", "fmt"))

	return g, nil
}
