package dashboard

import (
	"text/template"

	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/dashi/genny"
	"github.com/swiftcarrot/dashi/genny/gogen"
)

func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	name := opts.Name.Pascalize().Pluralize().ToLower().String()

	g.Transformer(genny.Replace("-pages-", "packages/dashboard/src/pages/"+name))
	g.Transformer(genny.Replace("-components-", "packages/components/src/"+name))

	err := g.Box(packr.New("scaffold:dashboard:templates", "../dashboard/templates"))
	if err != nil {
		return g, err
	}

	return g, nil
}
