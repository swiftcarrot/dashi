package generate

import (
	"context"
	"os/exec"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/flect"
	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/dashi/generators/graphql"
	"github.com/swiftcarrot/dashi/generators/migration"
	"github.com/swiftcarrot/dashi/generators/model"
	"github.com/swiftcarrot/dashi/generators/scaffold"
	"github.com/swiftcarrot/dashi/generators/scaffold/dashboard"
	"github.com/swiftcarrot/dashi/generators/scaffold/schema"
)

func getOptions(args []string) (*scaffold.Options, error) {
	if len(args) > 1 {
		name := args[0]
		as, err := attrs.ParseArgs(args[1:]...)
		if err != nil {
			return nil, err
		}

		opts := &scaffold.Options{
			Name:  flect.New(flect.Pascalize(name)).Singularize(),
			Attrs: as,
		}

		err = opts.Validate()
		if err != nil {
			return nil, err
		}

		return opts, nil
	}

	return nil, nil
}

var ScaffoldDashboardCmd = &cobra.Command{
	Use:     "scaffold:dashboard",
	Short:   "Generate scaffold for dashboard",
	Example: "dashi generate scaffold:dashboard Post title:string content:text",
	RunE: func(cmd *cobra.Command, args []string) error {
		gg := &genny.Group{}
		run := genny.WetRunner(context.Background())
		run.Logger = logger.New(logger.DebugLevel)

		opts, err := getOptions(args)

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

var ScaffoldCmd = &cobra.Command{
	Use:              "scaffold",
	Short:            "Generate scaffold for model",
	Example:          "dashi generate scaffold Post title:string content:text",
	TraverseChildren: true,
	Args:             cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		gg := &genny.Group{}
		run := genny.WetRunner(context.Background())
		run.Logger = logger.New(logger.DebugLevel)
		run.Logger.Infof("Generating new scaffold")

		opts, err := getOptions(args)
		if err != nil {
			return err
		}

		mops := &model.Options{
			Name:                   opts.Name,
			Attrs:                  opts.Attrs,
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

		schemaGen, err := schema.New(&scaffold.Options{
			Name:  opts.Name,
			Attrs: mops.Attrs,
		})
		if err != nil {
			return err
		}
		gg.Add(schemaGen)

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
