package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	usersGenerateHMACCmd = &cobra.Command{
		Use:     "generate-hmac",
		Short:   "Generate the base64-encoded HMAC signature of a user's email",
		Example: "mbctl users generate-hmac hana@magicbell.io",
		Aliases: []string{"gen-hmac"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			email := args[0]
			hmacEmail := api.GenerateUserEmailHMAC(email)

			if usersGenerateHMACCmdSimple {
				fmt.Println(hmacEmail)
			} else {
				logrus.Infof("HMAC of %s is %s", email, hmacEmail)
			}
		},
	}
	usersGenerateHMACCmdSimple bool
)

func init() {
	usersGenerateHMACCmd.Flags().BoolVarP(&usersGenerateHMACCmdSimple, "simple", "s", false, "Simple output, only the base64 hmac signature")

	usersCmd.AddCommand(usersGenerateHMACCmd)
}
