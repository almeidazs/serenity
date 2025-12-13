package cmd

import "github.com/spf13/cobra"

var CheckCmd = &cobra.Command{
	Use:   "check [path...]",
	Short: "Check code for issues",
	RunE:  check,
}

var checkWrite, checkUnsafe bool

func init() {
	rootCmd.AddCommand(CheckCmd)
	CheckCmd.Flags().BoolVarP(&checkWrite, "write", "w", false, "Write changes to files")
	CheckCmd.Flags().BoolVarP(&checkUnsafe, "unsafe", "u", false, "Apply unsafe fixes")
}

func check(cmd *cobra.Command, args []string) error {
	return nil
}
