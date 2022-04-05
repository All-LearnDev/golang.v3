package forms

import "time"

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

type FProject struct {
	ID             int       `json:"id" `
	Name           string    `json:"name" validate:"required" `
	Description    string    `json:"description" validate:"required"`
	CreatedAt      time.Time `json:"createdAt" `
	UpdatedAt      time.Time `json:"updatedAt" `
	CreateByUserId int
	UpdateByUserId int
	//Tasks       []Tasks   `json:"tasks" `
}

type Tasks struct {
	Id              int
	subject         string
	description     string
	ProjectId       int
	AssignToUserId  int
	DuaDate         time.Time
	CreatedByUserId int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
