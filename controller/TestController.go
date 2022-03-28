package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	modelmapper "gopkg.in/jeevatkm/go-model.v1"
)

type FBook struct {
	Id    int
	Name  string
	Price int
}

type Book struct {
	Id    int
	Name  string
	Price int
}

func Test(c echo.Context) error {

	fbook := FBook{Id: 12, Name: "Java ", Price: 1345}
	var book Book
	errs := modelmapper.Copy(&book, fbook)
	if errs == nil {
		fmt.Println(" Fname ", fbook.Name, " Book name ", book.Name)
		fmt.Println(" Fname ", fbook.Id, " Book name ", book.Id)
	}
	return c.JSON(http.StatusOK, book)
}
