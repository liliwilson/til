package cmd

import (
    "github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
    Use:     "publish",
    Short:   "publish my stuff",
    Run: func(cmd *cobra.Command, args []string) {
        Publish()
    },
}

func init() {
    rootCmd.AddCommand(publishCmd)
}
