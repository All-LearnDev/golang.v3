package bookRepository

import (
	"projects/configs"
	"projects/entitys"

	"gorm.io/gorm"
)

func InitializeData() {
	Connection := configs.GetConnection()
	Connection.Create(&entitys.Book{Name: "Java 8", AuthorName: "Dinh Nguyen"})
	Connection.Create(&entitys.Book{Name: "Java 14", AuthorName: "Karty Katy"})
	Connection.Create(&entitys.Book{Name: "Golang ", AuthorName: "Puktin "})

}

func FindById(Id int) entitys.Book {
	Connection := configs.GetConnection()
	var book entitys.Book
	Connection.First(&book, Id)
	return book
}
func AddBook(name string, authorName string) entitys.Book {
	Connection := configs.GetConnection()
	book := entitys.Book{Name: name, AuthorName: authorName}
	Connection.Create(&book)
	return book
}

func UpdateBook(ID int, name string, authorName string) entitys.Book {
	Connection := configs.GetConnection()
	book := FindById(ID)
	book.Name = name
	book.AuthorName = authorName
	Connection.Save(book)
	return book
}

func DeleteBookId(Id int) {
	Connection := configs.GetConnection()
	var book entitys.Book
	Connection.Delete(&book, Id)
}

func ListBook() []entitys.Book {
	Connection := configs.GetConnection()
	var lists []entitys.Book
	Connection.Find(&lists)
	return lists
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Paging(page int, pageSize int) []entitys.Book {
	var books []entitys.Book
	Connection := configs.GetConnection()
	Connection.Scopes(Paginate(page, pageSize)).Find(&books)
	return books
}
