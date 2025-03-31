package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-git",
	Short: "go-git is a simple git clone written in Go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from go-git!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
