package cmd

import "github.com/spf13/cobra"

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "Lists available projects",
    Long:  `Lists available projects`,
}

func init() {
    RootCmd.AddCommand(ListCmd)
}
