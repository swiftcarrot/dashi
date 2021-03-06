package cmd

import (
	"context"
	"go/build"
	"os"
	"path/filepath"

	"github.com/gobuffalo/logger"
	"github.com/spf13/cobra"
	"github.com/swiftcarrot/dashi/generators/new"
	"github.com/swiftcarrot/flect"
	"github.com/swiftcarrot/genny"
)

var FlagNewAPI bool

// TODO: use github.com/gobuffalo/envy
func getGopath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new dashi project",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		name := ""
		root := ""

		if len(args) == 0 {
			name = filepath.Base(pwd)
			root = pwd
		} else {
			name = args[0]
			root = filepath.Join(pwd, name)
		}

		run := genny.WetRunner(context.Background())
		run.Root = root
		run.Logger = logger.New(logger.DebugLevel)
		run.Logger.Infof("Creating new project in %s", root)

		gopath := getGopath()
		rel, err := filepath.Rel(gopath+"/src", root)
		if err != nil {
			return err
		}

		gg, err := new.New(&new.Options{
			Name:    flect.New(name),
			Package: rel,
			APIOnly: FlagNewAPI,
		})
		if err != nil {
			return err
		}

		run.WithGroup(gg)
		return run.Run()
	},
}

func init() {
	newCmd.Flags().BoolVar(&FlagNewAPI, "api", false, "API (GraphQL) only project")

	rootCmd.AddCommand(newCmd)
}
