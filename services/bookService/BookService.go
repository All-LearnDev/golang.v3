package bookService

import (
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/bookRepository"
)

func InitializeData() {
	bookRepository.InitializeData()
}

func ListBook() (error, []entitys.Book) {
	return bookRepository.ListBook()
}

func AddBook(fbook forms.FBook) (error, entitys.Book) {
	return bookRepository.AddBook(fbook.Name, fbook.AuthorName)
}

func UpdateBook(fbook forms.FBook) (error, entitys.Book) {
	return bookRepository.UpdateBook(fbook.ID, fbook.Name, fbook.AuthorName)

}

func FindById(id int) (entitys.Book, error) {
	book, error := bookRepository.FindById(id)
	if error == nil {
		return book, nil
	} else {
		return entitys.Book{}, error
	}

}

func DeleteBookId(Id int) error {

	return bookRepository.DeleteBookId(Id)
}

func Paging(page int, pageSize int) []entitys.Book {
	books := bookRepository.Paging(page, pageSize)
	return books

}
