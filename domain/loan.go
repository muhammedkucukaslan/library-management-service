package domain

import "time"

type Loan struct {
	ID           int       `json:"id"`
	BookID       int       `json:"book_id"`
	UserID       int       `json:"user_id"`
	StartedDate  time.Time `json:"started_date"`
	DueDate      time.Time `json:"due_date"`
	ReturnedDate time.Time `json:"returned_date"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewLoan(bookID, userID int) *Loan {

	now := time.Now()
	return &Loan{
		BookID:      bookID,
		UserID:      userID,
		StartedDate: now,
		DueDate:     now.AddDate(0, 0, 15),
		Status:      "BORROWED",
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
	l.ReturnedDate = time.Now()
}
