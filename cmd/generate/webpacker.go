package generate

import (
	"context"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/flect"
	"github.com/swiftcarrot/dashi/generators/webpacker"
)

var WebpackerCmd = &cobra.Command{
	Use:     "webpacker",
	Short:   "Generate a webpacker project under packages",
	Example: "dashi g webpacker app",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		g, err := webpacker.New(&webpacker.Options{
			Name: flect.New(args[0]),
		})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}
