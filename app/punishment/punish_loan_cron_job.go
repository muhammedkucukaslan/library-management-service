package punishment

import (
	"context"
	"fmt"
	"os"
)

type PunishLoanCronJob struct {
	repo Repository
	es   EmailService
}

func NewPunishLoanCronJob(repo Repository, es EmailService) *PunishLoanCronJob {
	return &PunishLoanCronJob{repo: repo, es: es}
}

type OverdueLoan struct {
	LoanID int
	UserID int
	BookID int
	Status string
}

func (p *PunishLoanCronJob) Execute() {

	ctx := context.Background()
	var userIDs []int

	// i wanted to use transaction here, so impleaded all operations to the repo function

	userIDs, err := p.repo.PunishOverdueLoans(ctx)
	if err != nil {
		fmt.Println("error while punishing overdue loans", err)
		return
	}

	var emails []string

	emails, err = p.repo.GetPunishedUserEmails(ctx, userIDs)
	if err != nil {
		fmt.Println("error while getting user emails", err)
		return
	}

	if emails != nil {
		err = p.es.SendPunishmentEmails(os.Getenv("EMAIL_SERVICE_DOMAIN"), emails, "Overdue Notice and Penalty Information", template)
		if err != nil {
			fmt.Println("error while sending punishment emails: ", err)
		}
	}
}

var template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Overdue Notice and Penalty Information</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 500px;
            margin: 0 auto;
            padding: 15px;
            color: #333;
        }
        h1 {
            color: #d9534f;
        }
        a {
            color: #2c5282;
            text-decoration: none;
        }
        .penalty {
            color: #d9534f;
            font-weight: bold;
        }
        .resolve-link {
            color: #337ab7;
        }
    </style>
</head>
<body>
    <h1>Overdue Notice and Penalty Information</h1>

	<p>Dear User,</p>
	<p>Our records indicate that you have overdue books that need to be returned. Please return them as soon as possible to avoid further penalties. You got 15 days punishment.</p>
	<p>If you have any questions or need assistance, please contact our support team.</p>

    <p>Sincerely,<br>- The Library Team</p>
</body>
</html>
`
