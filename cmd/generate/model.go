package generate

import (
	"context"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/model"
	"github.com/swiftcarrot/dashi/generators/scaffold/attrs"
)

var ModelCmd = &cobra.Command{
	Use:     "model",
	Short:   "Generate model",
	Example: "dashi generate model post title:string content:text",
	RunE: func(cmd *cobra.Command, args []string) error {
		as, err = attrs.ParseArgs(args[1:]...)
		if err != nil {
			return err
		}

		modelGen, err := model.New(&model.Options{
			Name:  flect.New(flect.Pascalize(n)).Singularize(),
			Attrs: as,
		})
		if err != nil {
			return err
		}

		run := genny.WetRunner(context.Background())
		run.With(g)
		return run.Run()
	},
}
