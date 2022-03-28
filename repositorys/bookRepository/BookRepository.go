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

func FindById(Id int) (entitys.Book, error) {
	Connection := configs.GetConnection()
	var book entitys.Book
	error := Connection.First(&book, Id).Error
	if error == nil {
		return book, nil
	} else {
		return entitys.Book{}, gorm.ErrRecordNotFound
	}
}
func AddBook(name string, authorName string) (error, entitys.Book) {
	Connection := configs.GetConnection()
	book := entitys.Book{Name: name, AuthorName: authorName}
	error := Connection.Create(&book)
	return error.Error, book
}

func UpdateBook(ID int, name string, authorName string) (error, entitys.Book) {
	Connection := configs.GetConnection()
	book, _ := FindById(ID)
	book.Name = name
	book.AuthorName = authorName
	error := Connection.Save(book)
	return error.Error, book

}

func DeleteBookId(Id int) error {
	Connection := configs.GetConnection()
	var book entitys.Book
	error := Connection.Delete(&book, Id)
	return error.Error
}

func ListBook() (error, []entitys.Book) {
	Connection := configs.GetConnection()
	var lists []entitys.Book
	error := Connection.Find(&lists)
	return error.Error, lists
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
