package mail

import (
	"fmt"
	"server/utils"

	"github.com/resend/resend-go/v2"
)

func SendMail(to []string, template MailTemplate) error {
	key, _key := utils.GetEnv("RESEND_API_KEY")

	if _key != nil {
		return fmt.Errorf("%w", _key)
	}

	client := resend.NewClient(key)

	params := &resend.SendEmailRequest{
		From: "onboarding@resend.dev",
		To: []string(to),
		Html: template.HTML,
		Subject: template.Subject,
		Text: template.Text,
	}

	_, _sent := client.Emails.Send(params)

	if _sent != nil {
		return fmt.Errorf("error sending email -> %w", _sent)
	}

	return nil
}
