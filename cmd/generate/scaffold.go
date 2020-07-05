package generate

import (
	"context"
	"os/exec"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/dashi/generators/graphql"
	"github.com/swiftcarrot/dashi/generators/migration"
	"github.com/swiftcarrot/dashi/generators/model"
	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/dashi/generators/scaffold/dashboard"
	"github.com/swiftcarrot/dashi/generators/scaffold/schema"
)

var ScaffoldCmd = &cobra.Command{
	Use:              "scaffold",
	Short:            "Generate scaffold for model",
	Example:          "dashi generate scaffold user id:int name:string desc:string",
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		gg := &genny.Group{}
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
		gg.Add(schemaGen)

		mops := &model.Options{
			Name:                   flect.New(flect.Pascalize(n)).Singularize(),
			Attrs:                  as,
			Path:                   "models",
			Package:                "models",
			TestPackage:            "models",
			Encoding:               "json",
			ForceDefaultID:         true,
			ForceDefaultTimestamps: true,
		}
		if err := mops.Validate(); err != nil {
			return err
		}

		modelGen, err := model.New(mops)
		if err != nil {
			return err
		}
		gg.Add(modelGen)

		// TODO add mysql support, remove hardcode postgres
		// migration attrs is from model opts which is validated and includes timestamp and default id column
		migrationGen, err := migration.New(&migration.Options{
			Dialect: "postgres",
			Name:    opts.Name,
			Time:    GetTime(),
			Attrs:   mops.Attrs,
		})
		if err != nil {
			return err
		}
		gg.Add(migrationGen)

		graphqlGen, err := graphql.New()
		if err != nil {
			return err
		}
		gg.Add(graphqlGen)

		dashboardGen, err := dashboard.New(opts)
		if err != nil {
			return err
		}
		gg.Add(dashboardGen)

		fmt := genny.New()
		fmt.RunFn(func(r *genny.Runner) error {
			r.Exec(exec.Command("make", "fmt"))
			return nil
		})
		gg.Add(fmt)

		run.WithGroup(gg)
		return run.Run()
	},
}
