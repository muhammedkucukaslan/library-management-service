package auth

type EmailService interface {
	SendWelcomeEmail(from, to, subject, html string) error
}
