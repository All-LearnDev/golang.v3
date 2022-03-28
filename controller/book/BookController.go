package controller

import (
	"net/http"
	"projects/forms"
	"projects/services/bookService"
	"projects/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func InitializeData(c echo.Context) error {
	bookService.InitializeData()
	return c.String(http.StatusOK, "Initialize ")
}

func ListBook(c echo.Context) error {
	error, list := bookService.ListBook()
	if error == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"result": true,
			"list":   list,
		})
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"result": false,
			"error":  error.Error(),
		})
	}
}

func AddBook(c echo.Context) error {
	var fbook forms.FBook
	c.Bind(&fbook)
	var validate = validator.New()
	err := validate.Struct(fbook)
	if err != nil {
		listError := utils.Validate(fbook)
		return c.JSON(http.StatusBadRequest, listError)
	} else {
		error, book := bookService.AddBook(fbook)
		if error == nil {
			return c.JSON(http.StatusOK, echo.Map{
				"result": true,
				"book":   book,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"result": false,
				"error":  error.Error(),
			})
		}

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
		return c.String(http.StatusOK, "Data Validated")
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
		error, book := bookService.UpdateBook(fbook)
		if error == nil {
			return c.JSON(http.StatusOK, echo.Map{
				"result": true,
				"book":   book,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"result": false,
				"error":  error.Error(),
			})
		}

	}

}

func FindById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	book, error := bookService.FindById(intVar)
	if error == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"result": true,
			"book":   book,
		})
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"result": false,
			"error":  error.Error(),
		})
	}

}

func DeleteBookId(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	error := bookService.DeleteBookId(intVar)
	if error == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"result": true,
		})

	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"result": false,
			"error":  error.Error(),
		})
	}
}

func Paging(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	pageSize, err := strconv.Atoi(c.Param("pageSize"))
	if err == nil {
		books := bookService.Paging(page, pageSize)
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
