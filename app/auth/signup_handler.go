package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

type SignupRequest struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Role       string    `json:"role"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	BirthDate  time.Time `json:"birth_date"`
}

type SignupResponse struct{}

type SignupHandler struct {
	repo Repository
	ts   TokenService
	cs   CookieService
	es   EmailService
}

func NewSignupHandler(repo Repository, ts TokenService, cs CookieService, es EmailService) *SignupHandler {
	return &SignupHandler{repo: repo, ts: ts, cs: cs, es: es}
}

func (h *SignupHandler) Handle(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	user := domain.NewUser(req.FirstName, req.SecondName, req.Role, req.Password, req.Email, req.Address, req.BirthDate)

	userID, err := h.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	go func() {
		err := h.es.SendWelcomeEmail(os.Getenv("EMAIL_SERVICE_DOMAIN"), user.Email, "Welcome to our library", createHTMLBody(user.FirstName))
		if err != nil {
			fmt.Println("error while sending welcome email: ", err)
		}
	}()

	token, err := h.ts.GenerateToken(userID, user.Role)
	if err != nil {
		return nil, err
	}

	err = h.cs.SetAuthCookies(ctx, token, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return &SignupResponse{}, nil
}

func createHTMLBody(firstName string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Welcome to Our Library</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			max-width: 600px;
			margin: 0 auto;
			padding: 20px;
			color: #333;
			background-color: #f9f9f9;
		}
		h1 {
			color: #2c5282;
		}
		p {
			line-height: 1.6;
		}
		a {
			color: #2c5282;
			text-decoration: none;
		}
		a:hover {
			text-decoration: underline;
		}
	</style>
</head>
<body>
	<h1>Welcome to Our Library</h1>

	<p>Dear %s,</p>
	<p>Welcome to our library! We hope you will gain a lot of useful knowledge and enjoy your time with us.</p>
	<p>Happy reading!</p>
	<p>- The Library Team</p>
</body>
</html>`, firstName)
}
