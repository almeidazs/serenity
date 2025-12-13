package cmd

import "github.com/spf13/cobra"

var FormatCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format files",
	RunE:  format,
}

var formatWrite, formatUnsafe bool

func init() {
	rootCmd.AddCommand(FormatCmd)
	FormatCmd.Flags().BoolVarP(&formatWrite, "write", "w", false, "Write changes to files")
	FormatCmd.Flags().BoolVarP(&formatUnsafe, "unsafe", "u", false, "Apply unsafe fixes")
}

func format(cmd *cobra.Command, args []string) error {
	return nil
}
