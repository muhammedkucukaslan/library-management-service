package domain

import "time"

type Author struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Bio        string    `json:"bio"`
	BirthDate  time.Time `json:"birth_date"`
}
