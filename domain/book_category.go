package domain

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type BookCategory struct {
	ID         int `json:"id"`
	CategoryID int `json:"category_id"`
	BookID     int `json:"book_id"`
}
