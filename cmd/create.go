package cmd

import "github.com/spf13/cobra"

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a project, occurrence, or note",
	Long:  `Creates a Grafeas project, occurrence, or note`,
}

func init() {
	RootCmd.AddCommand(CreateCmd)
}
