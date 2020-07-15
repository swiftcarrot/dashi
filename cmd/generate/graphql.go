package generate

import (
	"context"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/graphql"
)

var GraphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "generate a graphql server based on schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		run.Logger = logger.New(logger.DebugLevel)

		graphqlGen, err := graphql.New()
		if err != nil {
			return err
		}
		err = run.With(graphqlGen)
		if err != nil {
			return err
		}
		return run.Run()
	},
}
