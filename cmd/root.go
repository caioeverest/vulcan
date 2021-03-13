package cmd

import (
	"github.com/caioeverest/vulcan/infra/config"
	"github.com/caioeverest/vulcan/infra/console"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "vulcan [function]",
		PersistentPreRunE: prerun,
		SilenceUsage:      true,
	}
	c     console.Console
	debug bool
)

func Execute() error {
	c = console.Get()
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logs")
	if debug {
		c.SetDebugLevel()
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func prerun(cmd *cobra.Command, args []string) error {
	if debug {
		c.SetDebugLevel()
	}

	_ = config.Open()
	conf := config.Get()

	if conf.IsEmpty() {
		c.Warn("No configuration found, so we need to configure the app first")
		if err := configCMD.RunE(cmd, args); err != nil {
			return err
		}
	}
	return nil
}
