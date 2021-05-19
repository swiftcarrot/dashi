package new

import (
	"embed"
	"html/template"
	"os/exec"

	"github.com/swiftcarrot/dashi/generators/dashboard"
	"github.com/swiftcarrot/dashi/generators/packages"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/genny/gogen"
)

//go:embed templates
var templates embed.FS

func New(opts *Options) (*genny.Group, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	gg := &genny.Group{}

	g := genny.New()
	g.RunFn(func(r *genny.Runner) error {
		r.Exec(exec.Command("go", "mod", "init"))
		return nil
	})

	g.Templates(&templates)

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	g.Transformer(genny.Dot())
	gg.Add(g)

	if !opts.APIOnly {
		packages, err := packages.New(&packages.Options{
			Name: opts.Name,
		})
		if err != nil {
			return nil, err
		}
		gg.Add(packages)

		dashboard, err := dashboard.New(&dashboard.Options{
			Name: opts.Name,
		})
		if err != nil {
			return nil, err
		}
		gg.Add(dashboard)
	}

	make := genny.New()
	make.Command(exec.Command("make"))
	gg.Add(make)

	return gg, nil
}
