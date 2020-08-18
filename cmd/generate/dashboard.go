package generate

import (
	"context"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/dashboard"
	"github.com/swiftcarrot/dashi/generators/webpacker"
)

var DashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Generate dashboard under packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		gg := &genny.Group{}

		g1, err := webpacker.New(&webpacker.Options{})
		if err != nil {
			return err
		}

		g2, err := dashboard.New(&dashboard.Options{})
		if err != nil {
			return err
		}

		gg.Add(g1)
		gg.Add(g2)
		run.WithGroup(gg)
		return run.Run()
	},
}
