package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "til",
	Short: "post to my rss feed!",
	Long:  "til is a cli tool for publishing rss posts from markdown files",
	Run: func(cmd *cobra.Command, args []string) {
        Tui()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "oh no! something went wrong while executing til '%s'\n", err)
		os.Exit(1)
	}
}
