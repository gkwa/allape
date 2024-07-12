package cmd

import (
	"github.com/gkwa/allape/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Demonstrates different verbosity levels",
	Long:  `This command demonstrates the different verbosity levels of logging.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		core.Hello(logger)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
