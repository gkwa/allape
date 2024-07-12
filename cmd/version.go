package cmd

import (
	"github.com/gkwa/allape/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of allape",
	Long:  `All software has versions. This is allape's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		buildInfo := version.GetBuildInfo()
		logger.Info("Version info", "version", buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
