package generate

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/packages"
	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/genny"
)

var PackagesCmd = &cobra.Command{
	Use:     "packages",
	Short:   "Generate packages",
	Example: "dashi g packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := getName()
		if err != nil {
			return err
		}

		run := genny.WetRunner(context.Background())
		g, err := packages.New(&packages.Options{
			Name: flect.New(name),
		})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}
