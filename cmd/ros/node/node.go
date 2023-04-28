// Package node is used to hold the commands for the tool
package node

import (
	"sample/internal/logger"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var (
	packageName   string
	workspacePath string
	packageNode   string
	log           = logger.GetInstance()

	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "A set of node actions that can be done to a ros package",
		Long:  `Node allows user to interact with the nodes within a ros package.`,
	}
)

// New is use for exporting commands
func New() *cobra.Command {
	nodeCmd.AddCommand(newAddNode())
	return nodeCmd
}
