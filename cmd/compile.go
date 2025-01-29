package cmd

import (
    "github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
    Use:     "compile",
    // Aliases: []string{"c"},
    Short:   "compile markdown files into rss feed",
    Run: func(cmd *cobra.Command, args []string) {
        Compile()
    },
}

func init() {
    rootCmd.AddCommand(compileCmd)
}
