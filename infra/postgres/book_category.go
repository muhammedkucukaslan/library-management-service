package postgres

import (
	"context"

	bookcategory "github.com/muhammedkucukaslan/library-management-service/app/book_category"
)

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

func (r *Repository) GetCategories(ctx context.Context) (*bookcategory.GetBookCategoriesResponse, error) {
	query := `SELECT id, title FROM categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories bookcategory.GetBookCategoriesResponse
	for rows.Next() {
		var category bookcategory.Category
		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return &categories, nil
}
