package utils

import (
	"log"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
)

func SendCompletionEmail(email, foundationName string) error {
	client := mailjet.NewMailjetClient(os.Getenv("MAILJET_API_KEY"), os.Getenv("MAILJET_API_SECRET"))

	messages := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: os.Getenv("EMAIL_SENDER"),
				Name:  "Foundation Notifications",
			},
			To: &mailjet.RecipientsV31{
				{
					Email: email,
					Name:  foundationName,
				},
			},
			Subject:  "OrderList Completed!",
			TextPart: "Hello " + foundationName + ",\n\nYour orderlist has been successfully completed.",
			HTMLPart: "<h3>Hello " + foundationName + ",</h3><p>Your orderlist has been successfully completed.</p>",
		},
	}

	messagesBody := mailjet.MessagesV31{Info: messages}

	_, err := client.SendMailV31(&messagesBody)
	if err != nil {
		log.Println("Failed to send email:", err)
	}
	return err
}
