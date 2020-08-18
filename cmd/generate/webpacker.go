package generate

import (
	"context"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/webpacker"
)

var WebpackerInstallCmd = &cobra.Command{
	Use:     "webpacker:install",
	Short:   "Add webpacker support to project",
	Example: "dashi g webpacker:install",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		g, err := webpacker.New(&webpacker.Options{})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}
