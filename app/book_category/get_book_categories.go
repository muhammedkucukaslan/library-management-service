package bookcategory

import "context"

type GetBookCategoriesRequest struct{}

type GetBookCategoriesResponse []Category

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type GetBookCategoriesHandler struct {
	repo Repository
}

func NewGetBookCategoriesHandler(repo Repository) *GetBookCategoriesHandler {
	return &GetBookCategoriesHandler{repo: repo}
}

func (h *GetBookCategoriesHandler) Handle(ctx context.Context, req *GetBookCategoriesRequest) (*GetBookCategoriesResponse, error) {
	categories, err := h.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
