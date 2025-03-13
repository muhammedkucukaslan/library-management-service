package punishment

import (
	"context"
	"fmt"
)

type PunishLoanCronJob struct {
	repo Repository
}

func NewPunishLoanCronJob(repo Repository) *PunishLoanCronJob {
	return &PunishLoanCronJob{repo: repo}
}

type OverdueLoan struct {
	LoanID int
	UserID int
	BookID int
	Status string
}

func (p *PunishLoanCronJob) Execute() {

	fmt.Println("punishing overdue loans")

	ctx := context.Background()
	var userIDs []int

	// i wanted to use transaction here, so impleaded all operations to the repo function

	userIDs, err := p.repo.PunishOverdueLoans(ctx)
	if err != nil {
		fmt.Println("error while punishing overdue loans", err)
		return
	}

	fmt.Print(userIDs)

	// TODO send email to the user
}
