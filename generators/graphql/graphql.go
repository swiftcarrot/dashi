package graphql

import (
	"fmt"
	"os"

	"github.com/swiftcarrot/dashi/plugin/resolvergen"
	"github.com/swiftcarrot/genny"
	"github.com/swiftcarrot/gqlgen/api"
	"github.com/swiftcarrot/gqlgen/codegen/config"
	"github.com/swiftcarrot/gqlgen/plugin/modelgen"
)

func New() (*genny.Generator, error) {
	g := genny.New()
	g.RunFn(func(r *genny.Runner) error {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load config", err.Error())
			os.Exit(2)
		}

		return api.Generate(cfg,
			api.NoPlugins(),
			api.AddPlugin(modelgen.New()),
			api.AddPlugin(resolvergen.New()),
		)
	})

	return g, nil
}
