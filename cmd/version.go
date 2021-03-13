package cmd

import (
	"github.com/caioeverest/vulcan/infra/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCMD)
}

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "Version of the",
	Run: func(cmd *cobra.Command, args []string) {
		c.Info(config.Version)
	},
}
