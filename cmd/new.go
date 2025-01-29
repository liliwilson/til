package cmd

import (
    "github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
    Use:     "new",
    // Aliases: []string{"n"},
    Short:   "new markdown file",
    Long:    "create a new markdown file for a blog post",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        New(args[0])
    },
}

func init() {
    rootCmd.AddCommand(newCmd)
}
