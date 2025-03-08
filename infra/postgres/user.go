package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) (int, error) {

	err := r.db.QueryRowContext(ctx,
		"INSERT INTO users (first_name, second_name, role, password, email, address, birth_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.FirstName, user.SecondName, user.Role, user.Password, user.Email, user.Address, user.BirthDate).Scan(&user.ID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, errors.New("email already exists")
		}
		return 0, err
	}

	return user.ID, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", email)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User
	err := row.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Role, &user.Password, &user.Email, &user.Address, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User

	err := row.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Role, &user.Password, &user.Email, &user.Address, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id int) error {
	rslt, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	if rowsAffected, _ := rslt.RowsAffected(); rowsAffected == 0 {
		return errors.New("invalid id")
	}
	return nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *domain.User) error {
	rslt, err := r.db.ExecContext(ctx,
		"UPDATE users SET first_name = $1, second_name = $2, role = $3, email = $4, address = $5, birth_date = $6  WHERE id = $7",
		user.FirstName, user.SecondName, user.Role, user.Email, user.Address, user.BirthDate, user.ID)
	if err != nil {
		return err
	}
	if rowsAffected, _ := rslt.RowsAffected(); rowsAffected == 0 {
		return errors.New("invalid id")
	}
	return nil
}

func (r *Repository) CheckEmail(ctx context.Context, email string) error {
	row := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", email)
	if err := row.Err(); err != nil {
		return err
	}
	var id int
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return errors.New("email already exists")
		}
		fmt.Printf("error: %v\n", err)
		return err
	}
	return errors.New("email already exists")
}
