package postgres

import (
	"context"
	"errors"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) CreateAuthor(ctx context.Context, author *domain.Author) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO authors (first_name, second_name, bio, birth_date) VALUES ($1, $2, $3, $4)", author.FirstName, author.SecondName, author.Bio, author.BirthDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteAuthor(ctx context.Context, id int) error {
	rslt, err := r.db.ExecContext(ctx, "DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		return err
	}
	if rowsAffected, _ := rslt.RowsAffected(); rowsAffected == 0 {
		return errors.New("invalid id")
	}
	return nil
}
