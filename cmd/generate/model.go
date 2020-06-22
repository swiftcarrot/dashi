package generate

import (
	"context"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
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
		g, err := model.New(mops)
		if err != nil {
			return err
		}

		run.With(g)

		return run.Run()
	},
}
