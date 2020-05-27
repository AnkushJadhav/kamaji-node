package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AnkushJadhav/kamaji-node/logger"
	"github.com/AnkushJadhav/kamaji-node/pkg/app"

	"github.com/spf13/cobra"
)

var (
	// ConfigFile is the default config file path
	ConfigFile = ""
)

func genStartCmd() *cobra.Command {
	var cfgFile string

	startCmd := &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateFlagPaths(cfgFile); err != nil {
				logger.Errorf("Error while reading command line files : %v", err.Error())
				return err
			}

			go startShutdownHandler()
			if err := app.Start(cfgFile); err != nil {
				logger.Errorf("Error while starting app : %v", err.Error())
				return err
			}

			return app.Stop()
		},
		SilenceUsage: true,
	}

	startCmd.Flags().StringVarP(&cfgFile, "configfile", "c", ConfigFile, fmt.Sprintf("path of the config file for kamaji-node (defaults to %s)", ConfigFile))
	startCmd.MarkFlagRequired("configfile")

	return startCmd
}

func validateFlagPaths(cfgFile string) error {
	f, err := os.OpenFile(cfgFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}

func startShutdownHandler() {
	listener := make(chan os.Signal, 1)
	signal.Notify(listener, syscall.SIGINT)

	for {
		select {
		case <-listener:
			logger.Infoln("recieved SIGINT, stopping application")
			app.Stop()
		}
	}
}
