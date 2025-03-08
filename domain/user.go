package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Role       string    `json:"role"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	BirthDate  time.Time `json:"birth_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewUser(firstName, secondName, role, password, email, address string, birthDate time.Time) *User {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	password = string(hashedPassword)
	return &User{
		FirstName:  firstName,
		SecondName: secondName,
		Role:       role,
		Password:   password,
		Email:      email,
		Address:    address,
		BirthDate:  birthDate,
	}
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("invalid  credentials")
	}
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
