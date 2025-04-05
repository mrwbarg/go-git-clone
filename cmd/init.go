package cmd

import (
	"github.com/mrwbarg/go-git-clone/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new git repository in the given path. If no path is provided, the current directory will be used.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = "./"
		} else {
			path = args[0]
		}
		commands.InitCmd(path)
	},
}
