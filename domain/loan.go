package domain

import "time"

type Loan struct {
	ID         int       `json:"id"`
	BookID     int       `json:"book_id"`
	UserID     int       `json:"user_id"`
	StartedAt  time.Time `json:"started_at"`
	DueAt      time.Time `json:"due_at"`
	ReturnedAt time.Time `json:"returned_at"`
	Status     string    `json:"status"`
}

func NewLoan(bookID, userID int) *Loan {

	now := time.Now()
	return &Loan{
		BookID:    bookID,
		UserID:    userID,
		StartedAt: now,
		DueAt:     now.AddDate(0, 0, 15),
		Status:    "BORROWED",
	}
}

func (l *Loan) IsReturned() bool {
	return l.Status == "RETURNED"
}

func (l *Loan) IsOverdue() bool {
	return l.Status == "OVERDUE"
}

func (l *Loan) Return() {
	l.Status = "RETURNED"
	l.ReturnedAt = time.Now()
}
