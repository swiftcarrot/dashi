package cmd

import (
	"github.com/gobuffalo/packr/v2/jam"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack migration files with packr2",
	RunE: func(cmd *cobra.Command, args []string) error {
		return jam.Pack(jam.PackOptions{})
	},
}

func init() {
	rootCmd.AddCommand(packCmd)
}
