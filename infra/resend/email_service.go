package render

import "github.com/resend/resend-go/v2"

type ResendEmailService struct {
	apiKey string
}

func NewResendEmailService(apiKey string) *ResendEmailService {
	return &ResendEmailService{apiKey: apiKey}
}

func (s *ResendEmailService) SendWelcomeEmail(from, to, subject, html string) error {

	client := resend.NewClient(s.apiKey)

	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	_, err := client.Emails.Send(params)
	return err
}

func (s *ResendEmailService) SendPunishmentEmails(from string, emails []string, subject, html string) error {
	client := resend.NewClient(s.apiKey)

	params := &resend.SendEmailRequest{
		From:    from,
		To:      emails,
		Subject: subject,
		Html:    html,
	}

	_, err := client.Emails.Send(params)
	return err
}
