package cmd

import "github.com/spf13/cobra"

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "Group of commands related to Notifications",
}

func init() {
	rootCmd.AddCommand(notificationsCmd)
}
