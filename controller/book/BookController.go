package controller

import (
	"net/http"
	"projects/entitys"
	"projects/forms"
	services "projects/services/book"
	"projects/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func InitializeData(c echo.Context) error {
	services.InitializeData()
	return c.String(http.StatusOK, "Initialize ")
}

func ListBook(c echo.Context) error {
	list := services.ListBook()
	return c.JSON(http.StatusOK, list)
}

func AddBook(c echo.Context) error {
	var fbook forms.FBook
	var book entitys.Book
	c.Bind(&fbook)
	var validate = validator.New()
	err := validate.Struct(fbook)
	if err != nil {
		listError := utils.Validate(fbook)
		return c.JSON(http.StatusBadRequest, listError)
	} else {
		book = services.AddBook(fbook)
		return c.JSON(http.StatusOK, book)
	}
}

func ValidateBook(c echo.Context) error {
	var fbook forms.FBook
	c.Bind(&fbook)
	var validate = validator.New()
	err := validate.Struct(fbook)
	if err != nil {
		listError := utils.Validate(fbook)
		return c.JSON(http.StatusBadRequest, listError)

	} else {
		return c.String(http.StatusOK, "Check validate !")
	}
}

func UpdateBook(c echo.Context) error {
	var fbook forms.FBook
	c.Bind(&fbook)
	var validate = validator.New()
	err := validate.Struct(fbook)
	if err != nil {
		listError := utils.Validate(fbook)
		return c.JSON(http.StatusBadRequest, listError)

	} else {
		book := services.UpdateBook(fbook)
		return c.JSON(http.StatusOK, book)
	}

}

func FindById(c echo.Context) error {
	id := c.Param("id")
	intVar, err := strconv.Atoi(id)
	if err == nil {
		book := services.FindById(intVar)
		return c.JSON(http.StatusOK, book)
	} else {
		return c.String(http.StatusOK, "Record not found ")
	}
}

func Paging(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	pageSize, err := strconv.Atoi(c.Param("pageSize"))
	if err == nil {
		books := services.Paging(page, pageSize)
		return c.JSON(http.StatusOK, books)
	} else {
		return c.JSON(http.StatusOK, nil)
	}

}

func WriteExcelFile(c echo.Context) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	//list := services.ListBook()
	//book := list[i]
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	f.SaveAs("Book1.xlsx")
	return c.File("Book1.xlsx")
}
