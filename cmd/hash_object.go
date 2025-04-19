package cmd

import (
	"fmt"
	"slices"

	"github.com/mrwbarg/go-git-clone/internal/commands"
	"github.com/mrwbarg/go-git-clone/internal/models/object"
	"github.com/spf13/cobra"
)

var hashObjectWFlag bool

func init() {
	hashObjectCmd.Flags().BoolVarP(&hashObjectWFlag, "write", "w", false, "Write the object hash to the repo.")
	rootCmd.AddCommand(hashObjectCmd)
}

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object TYPE PATH",
	Short: "Creates an object hash and optionally writes it to the repository.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		typeArg := object.ObjectType(args[0])
		path := args[1]

		if !slices.Contains(object.ObjectTypes, typeArg) {
			return fmt.Errorf("invalid type: %s, allowed types are: %v", typeArg, object.ObjectTypes)
		}

		commands.HashObjectCmd(typeArg, path, hashObjectWFlag)
		return nil
	},
}
