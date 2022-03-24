package services

import (
	"projects/entitys"
	"projects/forms"
	repositorys "projects/repositorys/book"
)

func InitializeData() {
	repositorys.InitializeData()
}

func ListBook() []entitys.Book {
	list := repositorys.ListBook()
	return list
}

func AddBook(fbook forms.FBook) entitys.Book {
	book := repositorys.AddBook(fbook.Name, fbook.AuthorName)
	return book
}

func UpdateBook(fbook forms.FBook) entitys.Book {
	book := repositorys.UpdateBook(fbook.ID, fbook.Name, fbook.AuthorName)
	return book

}

func FindById(id int) entitys.Book {
	book := repositorys.FindById(id)
	return book
}

func Paging(page int, pageSize int) []entitys.Book {
	books := repositorys.Paging(page, pageSize)
	return books

}
