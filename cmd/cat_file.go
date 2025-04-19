package cmd

import (
	"github.com/mrwbarg/go-git-clone/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(catFileCmd)
}

var catFileCmd = &cobra.Command{
	Use:   "cat-file SHA",
	Short: "Prints the raw content of an object to STDOUT.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		objectArg := args[0]

		commands.CatFileCmd(objectArg)
		return nil
	},
}
