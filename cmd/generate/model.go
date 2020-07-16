package generate

import (
	"context"

	"github.com/swiftcarrot/dashi/generators/migration"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/flect"
	"github.com/swiftcarrot/dashi/generators/attrs"
	"github.com/swiftcarrot/dashi/generators/model"
)

var ModelCmd = &cobra.Command{
	Use:     "model",
	Short:   "Generate model",
	Example: "dashi generate model post title:string content:text",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) > 0 {
			name = args[0]
		}

		var (
			atts attrs.Attrs
			err  error
		)
		if len(args) > 1 {
			atts, err = attrs.ParseArgs(args[1:]...)
			if err != nil {
				return err
			}
		}

		run := genny.WetRunner(context.Background())

		mops := &model.Options{
			Name:                   flect.New(flect.Pascalize(name)).Singularize(),
			Attrs:                  atts,
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

		migrationGen, err := migration.New(&migration.Options{
			Dialect: "postgres",
			Name:    mops.Name,
			Time:    GetTime(),
			Attrs:   mops.Attrs,
		})
		if err != nil {
			return err
		}

		run.With(modelGen)
		run.With(migrationGen)

		return run.Run()
	},
}
