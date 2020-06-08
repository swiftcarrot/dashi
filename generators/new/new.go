package new

import (
	"html/template"
	"os/exec"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/dashboard"
	"github.com/swiftcarrot/dashi/generators/packages"
)

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

	g.Box(packr.New("dashi:generators:new", "../new/templates"))

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	g.Transformer(genny.Dot())
	gg.Add(g)

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

	make := genny.New()
	make.Command(exec.Command("make"))
	gg.Add(make)

	return gg, nil
}
