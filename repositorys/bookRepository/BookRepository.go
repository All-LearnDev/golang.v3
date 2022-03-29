package bookRepository

import (
	"projects/configs"
	"projects/entitys"

	"gorm.io/gorm"
)

var Connection = configs.GetConnection()

func InitializeData() {
	Connection.Create(&entitys.Book{Name: "Java 8", AuthorName: "Dinh Nguyen"})
	Connection.Create(&entitys.Book{Name: "Java 14", AuthorName: "Karty Katy"})
	Connection.Create(&entitys.Book{Name: "Golang ", AuthorName: "Puktin "})

}

func FindById(Id int) (entitys.Book, error) {
	var book entitys.Book
	error := Connection.First(&book, Id).Error
	if error == nil {
		return book, nil
	} else {
		return entitys.Book{}, gorm.ErrRecordNotFound
	}
}
func AddBook(name string, authorName string) (error, entitys.Book) {
	book := entitys.Book{Name: name, AuthorName: authorName}
	error := Connection.Create(&book)
	return error.Error, book
}

func UpdateBook(ID int, name string, authorName string) (error, entitys.Book) {
	book, _ := FindById(ID)
	book.Name = name
	book.AuthorName = authorName
	error := Connection.Save(book)
	return error.Error, book

}

func DeleteBookId(Id int) error {
	var book entitys.Book
	error := Connection.Delete(&book, Id)
	return error.Error
}

func ListBook() (error, []entitys.Book) {
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
	Connection.Scopes(Paginate(page, pageSize)).Find(&books)
	return books
}

func PagingV2(page int, pageSize int) *gorm.DB {

	var books []entitys.Book
	result := Connection.Find(&books)
	return result
}
