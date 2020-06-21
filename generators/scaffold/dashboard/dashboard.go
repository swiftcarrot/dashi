package dashboard

import (
	"text/template"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/genny/v2/gogen"
	"github.com/gobuffalo/packr/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
)

// dashboard will be generated based on graphql schema, including basic component/page and api sdk
// this one only handle newly created entity case(scaffold), we should also have one generator to sync between schema and js code
func New(opts *scaffold.Options) (*genny.Generator, error) {
	g := genny.New()
	//TODO generate api using graphql
	//cfg, err := config.LoadConfigFromDefaultLocations()
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
	//	os.Exit(2)
	//}
	//g.RunFn(func(r *genny.Runner) error {
	//	err = api.Generate(cfg, api.NoPlugins(),
	//		api.AddPlugin(dashboardgen.New(opts.Name.String())))
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, "dashboard", err.Error())
	//		os.Exit(3)
	//	}
	//	return nil
	//})

	data := map[string]interface{}{
		"opts": opts,
	}
	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)

	name := opts.Name.Dasherize().Pluralize().String()

	g.Transformer(genny.Replace("-pages-", "packages/dashboard/src/pages/"+name))
	g.Transformer(genny.Replace("-components-", "packages/dashboard/src/components/"+name))

	err := g.Box(packr.New("scaffold:dashboard:templates", "../dashboard/templates"))
	if err != nil {
		return g, err
	}

	return g, nil
}
