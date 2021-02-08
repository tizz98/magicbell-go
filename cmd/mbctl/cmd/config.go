package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Group of config related commands",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
