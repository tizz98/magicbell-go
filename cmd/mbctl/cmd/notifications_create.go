package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	magicbell "github.com/tizz98/magicbell-go"
)

type notificationsCreateOptions struct {
	Title      string
	Recipients []string
	Content    string
	ActionURL  string
	Category   string
}

var (
	notificationsCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"send"},
		Short:   "Send a notification to one or more users.",
		RunE: func(cmd *cobra.Command, args []string) error {
			recipients := make([]magicbell.NotificationRecipient, len(notificationCreateOpts.Recipients))

			for i, r := range notificationCreateOpts.Recipients {
				if strings.ContainsRune(r, '@') {
					recipients[i] = magicbell.NotificationRecipient{Email: r}
				} else {
					recipients[i] = magicbell.NotificationRecipient{ExternalID: r}
				}
			}

			notification, err := api.CreateNotificationC(cmd.Context(), magicbell.CreateNotificationRequest{
				Title:            notificationCreateOpts.Title,
				Recipients:       recipients,
				Content:          notificationCreateOpts.Content,
				CustomAttributes: nil,
				ActionURL:        notificationCreateOpts.ActionURL,
				Category:         notificationCreateOpts.Category,
			})

			if err != nil {
				return err
			}

			logrus.Infof("Created notification %s", notification.ID)
			return nil
		},
	}

	notificationCreateOpts = &notificationsCreateOptions{}
)

func init() {
	notificationsCreateCmd.Flags().StringVar(&notificationCreateOpts.Title, "title", "", "The title of the notification")
	notificationsCreateCmd.Flags().StringSliceVar(&notificationCreateOpts.Recipients, "recipients", nil, "Comma separated emails or external ids of users who should receive this notification")
	notificationsCreateCmd.Flags().StringVar(&notificationCreateOpts.Content, "content", "", "The content of the notification")
	notificationsCreateCmd.Flags().StringVar(&notificationCreateOpts.ActionURL, "action-url", "", "The URL to redirect to when clicking the notification")
	notificationsCreateCmd.Flags().StringVar(&notificationCreateOpts.Category, "category", "", "The category of the notification")

	_ = notificationsCreateCmd.MarkFlagRequired("title")
	_ = notificationsCreateCmd.MarkFlagRequired("recipients")

	notificationsCmd.AddCommand(notificationsCreateCmd)
}
