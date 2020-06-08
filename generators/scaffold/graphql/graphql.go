package graphql

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/dashi/generators/scaffold/graphql/plugins/modelgen"
	"github.com/swiftcarrot/dashi/generators/scaffold/graphql/plugins/resolvergen"
)

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag += ` db:"` + flect.Underscore(field.Name) + `"`
		}
	}

	return b
}
func New(opts *scaffold.Options) (*genny.Generator, error) {

	g := genny.New()
	g.RunFn(func(r *genny.Runner) error {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
			os.Exit(2)
		}
		err = api.Generate(cfg, api.NoPlugins(),
			api.AddPlugin(&modelgen.Plugin{MutateHook: mutateHook}),
			api.AddPlugin(resolvergen.New()))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(3)
		}
		return nil
	})

	return g, nil
}
