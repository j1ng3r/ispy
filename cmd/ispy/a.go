package ispy

import (
	"github.com/j1ng3r/ispy/pkg/ispy"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:     "a",
	Aliases: []string{"a"},
	Short:   "adds a",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ispy.Append_a()
	},
}

func init() {
	rootCmd.AddCommand(aCmd)
}
