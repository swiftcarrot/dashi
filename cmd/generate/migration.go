package generate

import (
	"context"
	"time"

	"github.com/gobuffalo/genny/v2"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/migration"
	"github.com/swiftcarrot/flect"
)

func GetTime() string {
	t := time.Now()

	return t.Format("20060102150405")
}

var MigrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Generate new migration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		time := GetTime()

		run := genny.WetRunner(context.Background())
		// TODO remove hardcode postgres
		g, err := migration.New(&migration.Options{
			Name:    flect.New(name),
			Time:    time,
			Dialect: "postgres",
		})
		if err != nil {
			return err
		}
		run.With(g)
		return run.Run()
	},
}
