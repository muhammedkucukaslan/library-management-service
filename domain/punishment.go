package domain

import "time"

// TODO we may add more fields to punishment such as status

type Punishment struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	PunisherID int       `json:"punisher_id"`
	Reason     string    `json:"reason"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewPunishment(userID, punisherID int, reason string, endDate time.Time) *Punishment {
	return &Punishment{
		UserID:     userID,
		PunisherID: punisherID,
		Reason:     reason,
		EndDate:    endDate,
		CreatedAt:  time.Now(),
	}
}
