// Package node is used to hold the commands for the tool
package node

import (
	"sample/internal/logger"
	"sample/pkg/ros"

	"github.com/spf13/cobra"
)

// addCmd represents the nodeadd command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a node into an existing package",
	Long: `Add allows the user to create a new node in a given ros package. If the package
does not exist, it will throw an error. If the command already exists, it will throw an error.
For example:
sample ros node add --package-name mynewrospackage --package-path workspace --node mynode
sample ros node add -n mynewrospackage --node mynode`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log.Infof("Adding node %s to package %s", packageNode, packageName)
		ros.NewRosService(log)
	},
}

func newAddNode() *cobra.Command {
	addCmd.Flags().StringVarP(&packageName, "package-name", "n", "", "[required] package config")
	addCmd.MarkFlagRequired("package-name")
	addCmd.Flags().StringVar(&packageNode, "node", "", "[required] node name to be added to the package")
	addCmd.MarkFlagRequired("node")
	addCmd.Flags().StringVarP(&workspacePath, "workspace-path", "w", "workspace", "path to the workspace.")
	return addCmd
}
