package entitys

import "time"

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Image     string
}

type Images struct {
	Id        int `gorm:"primaryKey"`
	Filename  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshToken struct {
	UserId    int `gorm:"primaryKey"`
	UserName  string
	ExpiresAt int64
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Project struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreateByUserId int
	UpdateByUserId int
	//Tasks       []Tasks   `json:"tasks" `
}
type Tasks struct {
	Id              int `gorm:"primaryKey"`
	subject         string
	description     string
	ProjectId       int
	AssignToUserId  int
	DuaDate         time.Time
	CreatedByUserId int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
