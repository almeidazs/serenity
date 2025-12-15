package cmd

import (
	"fmt"
	"os/exec"

	"github.com/almeidazs/gowther/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes gowther, creating a json configuration file", // TODO: change this later
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

// NOTE: quais flags posso receber aq?
func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() error {
	var cmd *exec.Cmd

	path, err := config.GetPath()
	if err != nil {
		return err
	}

	exists, err := config.CheckHasConfigFile(path)
	if err != nil {
		return fmt.Errorf("error to find config file: %w", err)
	}

	if !exists {
		if err := config.CreateConfigFile(path); err != nil {
			return err
		}
	}

	return cmd.Start()
}
