// Package cmd contains the commands for the sample tool.
package cmd

import (
	"fmt"
	"os"
	"sample/cmd/ros"
	"sample/internal/logger"

	"github.com/spf13/cobra"
)

var (
	hash    string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "sample",
		Short: "sample",
		Long: `



	██╗███╗   ███╗██████╗
	██║████╗ ████║██╔══██╗
	██║██╔████╔██║██████╔╝
	██║██║╚██╔╝██║██╔═══╝
	██║██║ ╚═╝ ██║██║
	╚═╝╚═╝     ╚═╝╚═╝
												`,
	}
)

// Execute adds all child commands to the root command.
func Execute(version, commit string) {
	rootCmd.Version = version
	hash = commit

	setVersion()

	if err := rootCmd.Execute(); err != nil {
		logger.GetInstance().Error(err)
		os.Exit(1)
	}
}

func setVersion() {
	template := fmt.Sprintf("sample version: %s commit: %s \n", rootCmd.Version, hash)
	rootCmd.SetVersionTemplate(template)
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "set logging level to verbose")
	rootCmd.AddCommand(ros.New())
}
