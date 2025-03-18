package book

import "context"

type GetBooksRequest struct {
	AuthorID   int `json:"author_id" query:"author_id"`
	CategoryId int `json:"category_id" query:"category_id"`
	Limit      int `json:"limit" query:"limit"`
	Page       int `json:"page" query:"page"`
}

type GetBooksResponse []Book

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
}

type GetBooksHandler struct {
	repo Repository
}

func NewGetBooksHandler(repo Repository) *GetBooksHandler {
	return &GetBooksHandler{repo: repo}
}

func (h *GetBooksHandler) Handle(ctx context.Context, req *GetBooksRequest) (*GetBooksResponse, error) {

	if req.AuthorID != 0 {
		books, err := h.repo.GetBooksByAuthor(ctx, req.AuthorID)
		if err != nil {
			return nil, err
		}
		return books, nil
	}

	if req.CategoryId != 0 {
		books, err := h.repo.GetBooksByCategory(ctx, req.CategoryId)
		if err != nil {
			return nil, err
		}
		return books, nil
	}

	books, err := h.repo.GetAllBooks(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}
