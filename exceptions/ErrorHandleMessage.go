package exceptions

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const FILE_NOT_FOUND string = "File Not Found"
const RECORD_NOT_FOUND string = "Record Not Found"
const IN_VALID_TOKEN string = "Invalid token "
const DATABASE_CONNECTION_ERROR string = "Database connection error"
const EMAIL_EXIST string = "Email exits in DB"
const UN_AUTHORIZED = "Un Authorized"
const IN_VALID_USERNAME_PASSWORD = "User name or password is incorrect"
const VALIDATION_EXCEPTION string = "Validation exception"

func RecordNotFoundException(storeErr error, c echo.Context) error {

	return c.JSON(http.StatusBadRequest, echo.Map{
		"result":  false,
		"message": RECORD_NOT_FOUND,
		"error":   storeErr.Error(),
	})

}

func EmailExistsDBException(c echo.Context) error {

	return c.JSON(http.StatusBadRequest, echo.Map{
		"result":  false,
		"message": EMAIL_EXIST,
	})

}
func InValidTokenException(c echo.Context) error {

	return c.JSON(http.StatusBadRequest, echo.Map{
		"result":  false,
		"message": IN_VALID_TOKEN,
	})

}

func UnauthorizedException(c echo.Context) error {

	return c.JSON(http.StatusForbidden, echo.Map{
		"result":  false,
		"message": UN_AUTHORIZED,
	})

}

func IncorrectUserNamePasswordException(c echo.Context) error {

	return c.JSON(http.StatusForbidden, echo.Map{
		"result":  false,
		"message": IN_VALID_USERNAME_PASSWORD,
	})

}

func DatabaseConnectionException(err error, c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"result":  false,
		"message": DATABASE_CONNECTION_ERROR,
		"error":   err.Error(),
	})

}

func StoreFileException(storeErr error, c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"result":  false,
		"message": FILE_NOT_FOUND,
		"error":   storeErr.Error(),
	})

}

func ValidationFieldException(error []FieldError, c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"result":  false,
		"message": VALIDATION_EXCEPTION,
		"error":   error,
	})

}
func Validate(form interface{}) []FieldError {

	var validate = validator.New()
	err := validate.Struct(form)
	var listError []FieldError
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage FieldError
			errorMessage.Field = err.Field()
			errorMessage.Error = err.Error()
			listError = append(listError, errorMessage)
		}
	}
	return listError
}
