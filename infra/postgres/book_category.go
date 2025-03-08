package postgres

import "context"

func (r *Repository) CreateCategory(ctx context.Context, title string) error {
	query := `INSERT INTO categories (title) VALUES ($1)`
	_, err := r.db.ExecContext(ctx, query, title)
	return err
}

func (r *Repository) DeleteCategory(ctx context.Context, id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
