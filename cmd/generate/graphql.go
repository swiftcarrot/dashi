package generate

import (
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var GraphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "generate a graphql server based on schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if os.IsNotExist(errors.Cause(err)) {
			cfg = config.DefaultConfig()
		} else if err != nil {
			return err
		}

		if err = api.Generate(cfg); err != nil {
			return err
		}

		return nil
	},
}
