package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	magicbell "github.com/tizz98/magicbell-go"
)

type notificationsCreateOptions struct {
	Title            string
	Recipients       []string
	Content          string
	ActionURL        string
	Category         string
	CustomAttributes []string // Key=Value
}

func (o notificationsCreateOptions) getNotificationRecipients() (recipients []magicbell.NotificationRecipient) {
	recipients = make([]magicbell.NotificationRecipient, len(notificationCreateOpts.Recipients))

	for i, r := range notificationCreateOpts.Recipients {
		if strings.ContainsRune(r, '@') {
			recipients[i] = magicbell.NotificationRecipient{Email: r}
		} else {
			recipients[i] = magicbell.NotificationRecipient{ExternalID: r}
		}
	}

	return
}

func (o notificationsCreateOptions) getCustomAttributes() magicbell.CustomAttributes {
	if o.CustomAttributes == nil {
		return nil
	}

	attrs := magicbell.CustomAttributes{}
	for _, rawAttr := range o.CustomAttributes {
		values := strings.SplitN(rawAttr, "=", 2)
		if len(values) != 2 {
			panic(fmt.Sprintf("expected '=' in custom attribute, got %s", rawAttr))
		}

		attrs[values[0]] = values[1]
	}

	return attrs
}

var (
	notificationsCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"send"},
		Short:   "Send a notification to one or more users.",
		RunE: func(cmd *cobra.Command, args []string) error {
			notification, err := api.CreateNotificationC(cmd.Context(), magicbell.CreateNotificationRequest{
				Title:            notificationCreateOpts.Title,
				Recipients:       notificationCreateOpts.getNotificationRecipients(),
				Content:          notificationCreateOpts.Content,
				CustomAttributes: notificationCreateOpts.getCustomAttributes(),
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
	notificationsCreateCmd.Flags().StringArrayVar(&notificationCreateOpts.CustomAttributes, "custom-attribute", nil, "A list of custom attributes in the Key=Value format")

	_ = notificationsCreateCmd.MarkFlagRequired("title")
	_ = notificationsCreateCmd.MarkFlagRequired("recipients")

	notificationsCmd.AddCommand(notificationsCreateCmd)
}
