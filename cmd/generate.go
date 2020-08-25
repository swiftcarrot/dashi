package cmd

import (
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/cmd/generate"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generates dashboard, scaffold, migration, ...",
	Long:    "Create new project",
}

func init() {
	generateCmd.AddCommand(generate.DashboardCmd)
	generateCmd.AddCommand(generate.ScaffoldCmd)
	generateCmd.AddCommand(generate.ScaffoldDashboardCmd)
	generateCmd.AddCommand(generate.MigrationCmd)
	generateCmd.AddCommand(generate.ModelCmd)
	generateCmd.AddCommand(generate.GraphqlCmd)
	generateCmd.AddCommand(generate.PackagesCmd)

	rootCmd.AddCommand(generateCmd)
}
