package book

type GetBooksRequest struct {
	AuthorID   int `json:"author_id"`
	CategoryId int `json:"category_id"`
}
