package domain

import (
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	AuthorID    int       `json:"author_id"`
	ISBN        string    `json:"isbn"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	Publisher   string    `json:"publisher"`
	Edition     int       `json:"edition"`
	Stock       BookStock `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookStock struct {
	ID        int       `json:"id"`
	BookID    int       `json:"book_id"`
	Total     int       `json:"total"`
	Available int       `json:"available"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBook(title, ISBN, description, publisher string, authorID, edition int) *Book {
	return &Book{
		Title:       title,
		ISBN:        ISBN,
		Description: description,
		AuthorID:    authorID,
		Publisher:   publisher,
		Edition:     edition,
	}
}

func NewStock(bookID, total int) *BookStock {
	return &BookStock{
		BookID:    bookID,
		Total:     total,
		Available: total,
	}
}
