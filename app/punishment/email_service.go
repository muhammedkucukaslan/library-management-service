package punishment

type EmailService interface {
	SendPunishmentEmails(from string, emails []string, subject, html string) error
}
