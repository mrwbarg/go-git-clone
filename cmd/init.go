package cmd

import (
	"github.com/mrwbarg/go-git-clone/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new git repository",
	Run: func(cmd *cobra.Command, args []string) {
		commands.InitCmd()
	},
}
