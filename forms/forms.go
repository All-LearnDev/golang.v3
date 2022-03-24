package forms

type FBook struct {
	ID         int    `json:"id" `
	Name       string `json:"name" validate:"required"`
	AuthorName string `json:"authorName" validate:"required" `
}

type FProject struct {
	ID       int    `json:"id" `
	Name     string `json:"name" `
	Customer string `json:"customer" `
	//developers []FDeveloper `json:"developers"`
}

type FDeveloper struct {
	Id   int        `json:"id" `
	Name string     `json:"name" `
	Age  int        `json:"age" `
	List []FProject `json:"list"`
}
