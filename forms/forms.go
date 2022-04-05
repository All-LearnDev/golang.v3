package forms

type FUser struct {
	ID       int
	Name     string ` validate:"required"`
	Email    string ` validate:"required" `
	Password string ` validate:"required" `
	Image    string ` validate:"required"`
}

type FLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

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
	Id       int        `json:"id" `
	Name     string     `json:"name" `
	Age      int        `json:"age" `
	Projects []FProject `json:"projects"`
}
