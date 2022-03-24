package entitys

type Book struct {
	//gorm.Model
	ID         int
	Name       string
	AuthorName string
}
type Author struct {
	//gorm.Model
	ID      int
	Name    string
	Address string
}

type User struct {
	//gorm.Model
	Id     int
	Name   string
	Email  string
	Images []Images
}

type Images struct {
	//gorm.Model
	Id       int
	Filename string
	UserID   int
}

type JUser struct {
	Id       int
	Name     string
	Email    string
	Password string `json:"-"`
}

type RefreshToken struct {
	UserId    int
	UserName  string
	ExpiresAt int64
	Token     string
}
type Project struct {
	ID         int
	Name       string
	price      uint64
	developers []*Developer `gorm:"many2many:dev_project;"`
}

type Developer struct {
	Id       int
	Name     string
	projects []*Project `gorm:"many2many:dev_project;"`
}
