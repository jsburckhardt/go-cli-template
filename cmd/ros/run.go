// Package ros is the root command for sample command.
package ros

import (
	"sample/internal/logger"
	"sample/pkg/ros"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "helper for running ros packages",
	Long: `This commands should be use to run ros packages given a package name, package path
and a target script. For example:

sample ros run --workspace-path workspace -n talkerlistener --node talker
sample ros run -p workspace -n talkerlistener --node talker
sample ros run -n talkerlistener --node talker`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log.Infof("running script %s in path %s", scriptName, packagePath)
		ros := ros.NewRosService(log)
		err := ros.Run(packageName, workspacePath, packageNode)
		if err != nil {
			log.Errorf("Unable to run script: %v", err)
		}
	},
}

func newCmdRosRun() *cobra.Command {
	runCmd.Flags().StringVarP(&packageName, "package-name", "n", "", "[required] package name")
	runCmd.MarkFlagRequired("package-name")
	runCmd.Flags().StringVarP(&workspacePath, "workspace-path", "w", "workspace", "path to the workspace.")
	runCmd.Flags().StringVar(&packageNode, "node", "", "[required] run a node in the package.")
	runCmd.MarkFlagRequired("node")
	return runCmd
}
