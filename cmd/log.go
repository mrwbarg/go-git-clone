package cmd

import (
	"github.com/mrwbarg/go-git-clone/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

var logCmd = &cobra.Command{
	Use:   "log COMMIT",
	Short: "Logs the commit history starting from the given commit.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sha := args[0]
		commands.LogCmd(sha)
	},
}
