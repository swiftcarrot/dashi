package graphql

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/gobuffalo/genny/v2"
	"github.com/swiftcarrot/dashi/generators/graphql/plugins/modelgen"
	"github.com/swiftcarrot/dashi/generators/graphql/plugins/resolvergen"
)

func New() (*genny.Generator, error) {
	g := genny.New()
	g.RunFn(func(r *genny.Runner) error {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
			os.Exit(2)
		}
		err = api.Generate(cfg, api.NoPlugins(),
			api.AddPlugin(modelgen.New()),
			api.AddPlugin(resolvergen.New()))
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(3)
		}
		return nil
	})

	return g, nil
}
