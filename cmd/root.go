package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "serenity",
	SilenceUsage: true,
	SilenceErrors: true,
	Short: " Serenity, an agressive and ultra fast Go linter with no noise",
}

func Exec() {
	cobra.CheckErr(rootCmd.Execute())
}
