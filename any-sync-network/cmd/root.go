package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var defaultsFlag bool
var outputDirFlag string

var rootCmd = &cobra.Command{
	Use:   "anyconf",
	Short: "Configuration builder for Any-Sync nodes.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	create.Flags().BoolVar(&defaultsFlag, "defaults", false, "generate configuration files using default parameters")
	create.Flags().StringVar(&outputDirFlag, "output-dir", ".", "directory where the config files are generated")
	rootCmd.AddCommand(create)
}
