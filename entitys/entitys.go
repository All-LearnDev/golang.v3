package entitys

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Image     string
}

type Images struct {
	Id        int
	Filename  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshToken struct {
	UserId    int
	UserName  string
	ExpiresAt int64
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Project struct {
	ID        int
	Name      string
	Customer  string
	Tasks     []Tasks `gorm:"many2many:dev_project;"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
