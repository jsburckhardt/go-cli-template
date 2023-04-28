// Package ros is the root command for sample command.
package ros

import (
	"sample/cmd/ros/node"
	"sample/internal/logger"

	"github.com/spf13/cobra"
)

// rosCmd represents the ros command
var (
	packagePath           string
	scriptName            string
	packageName           string
	workspacePath         string
	additionalArgs        string
	registry              string
	containerTarget       string
	containerVersion      string
	registryInsecure      bool
	isCoverageEnabled     bool
	pushContainer         bool
	amdOnly               bool
	armOnly               bool
	relativeWorkSpacePath string
	testResultPath        string
	packageNode           string
	log                   = logger.GetInstance()

	rosCmd = &cobra.Command{
		Use:   "ros",
		Short: "ros package actions",
		Long:  `ros contains a set of actions that can be done with ros packages.`,
	}
)

// New is use for exporting commands
func New() *cobra.Command {
	rosCmd.AddCommand(newCmdRosRun())
	rosCmd.AddCommand(node.New())
	return rosCmd
}
