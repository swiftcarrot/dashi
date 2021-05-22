package new

import (
	"embed"
	"html/template"
	"os/exec"

	"github.com/swiftcarrot/dashi/generators/dashboard"
	"github.com/swiftcarrot/dashi/generators/packages"
	"github.com/swiftcarrot/dashi/generators/server"
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

	g.Command(exec.Command("go", "mod", "init"))
	// TODO: lock package version
	g.Command(exec.Command("go", "get", "github.com/swiftcarrot/dashi"))
	g.Command(exec.Command("go", "get", "github.com/swiftcarrot/gqlgen"))

	g.Templates(&templates)

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	g.Transformer(genny.Replace("$dot$", "."))
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
	make.Command(exec.Command("make", "graphql"))
	gg.Add(make)

	server, err := server.New(&server.Options{
		Name:    opts.Name,
		Package: opts.Package,
	})
	if err != nil {
		return nil, err
	}
	server.Command(exec.Command("go", "get"))
	gg.Add(server)

	return gg, nil
}
