package bookService

import (
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/bookRepository"
)

func InitializeData() {
	bookRepository.InitializeData()
}

func ListBook() []entitys.Book {
	list := bookRepository.ListBook()
	return list
}

func AddBook(fbook forms.FBook) entitys.Book {
	book := bookRepository.AddBook(fbook.Name, fbook.AuthorName)
	return book
}

func UpdateBook(fbook forms.FBook) entitys.Book {
	book := bookRepository.UpdateBook(fbook.ID, fbook.Name, fbook.AuthorName)
	return book

}

func FindById(id int) entitys.Book {
	book := bookRepository.FindById(id)
	return book
}

func Paging(page int, pageSize int) []entitys.Book {
	books := bookRepository.Paging(page, pageSize)
	return books

}
