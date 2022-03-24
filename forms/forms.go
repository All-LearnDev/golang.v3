package forms

type FBook struct {
	ID         int    `json:"id" `
	Name       string `json:"name" validate:"required"`
	AuthorName string `json:"authorName" validate:"required" `
}
