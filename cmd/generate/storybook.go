package generate

import (
	"context"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/storybook"
)

var StorybookInstallCmd = &cobra.Command{
	Use:   "storybook:install",
	Short: "Add storybook suppport to project",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		run.Logger = logger.New(logger.DebugLevel)

		g, err := storybook.New(&storybook.Options{})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}
