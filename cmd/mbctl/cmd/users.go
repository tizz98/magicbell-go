package cmd

import "github.com/spf13/cobra"

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Commands related to Users",
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
