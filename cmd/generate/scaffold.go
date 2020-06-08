package generate

import (
	"context"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/dashi/generators/scaffold/attrs"
	"github.com/swiftcarrot/dashi/generators/scaffold/graphql"
	"github.com/swiftcarrot/dashi/generators/scaffold/migration"
	"github.com/swiftcarrot/dashi/generators/scaffold/schema"
)

var ScaffoldCmd = &cobra.Command{
	Use:              "scaffold",
	Short:            "Generate scaffold for model",
	Example:          "dashi generate scaffold user id:int name:string desc:string",
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		n := args[0]
		var (
			as  attrs.Attrs
			err error
		)
		if len(args) > 1 {
			as, err = attrs.ParseArgs(args[1:]...)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		run := genny.WetRunner(context.Background())
		run.Logger = logger.New(logger.DebugLevel)
		run.Logger.Infof("Creating new scaffold")
		var opts = &scaffold.Options{
			Name:  flect.New(flect.Pascalize(n)).Singularize(),
			Attrs: as,
		}
		if err := opts.Validate(); err != nil {
			return err
		}
		schemaGen, err := schema.New(opts)
		if err != nil {
			return err
		}
		graphqlGen, err := graphql.New(opts)
		if err != nil {
			return err
		}
		migrationGen, err := migration.New(opts)
		if err != nil {
			return err
		}
		//modelGen, err := model.New(opts)
		//if err != nil {
		//	return err
		//}

		err = run.With(schemaGen)
		err = run.With(migrationGen)
		//err = run.With(modelGen)
		err = run.With(graphqlGen)

		return run.Run()
	},
}
