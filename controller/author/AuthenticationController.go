package controller

import (
	"net/http"
	"projects/entitys"
	"projects/exceptions"
	"projects/forms"
	"projects/services/authorService"
	"projects/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

/**
 * User login.
 *
 * @param  email
 * @param  password
 * @return  login success or fail
 */
func Login(c echo.Context) error {

	var fuser forms.FLogin
	c.Bind(&fuser)
	var validate = validator.New()
	err := validate.Struct(fuser)
	if err != nil {
		listError := exceptions.Validate(fuser)
		return exceptions.ValidationFieldException(listError, c)
	} else {
		// Throws unauthorized error
		var user entitys.User
		_, user = authorService.GetUserByEmail(fuser.Email)
		if (fuser.Email != user.Email) || (utils.CheckPasswordHash(fuser.Password, user.Password) != true) {
			return exceptions.IncorrectUserNamePasswordException(c)
		}
		// Generate access_token
		accessToken := utils.GenerateJWT(user.Id, user.Name)
		// Generate refreshToken
		refreshToken := utils.GenerateRefreshToken(user.Id, user.Name)
		// Save refreshToken to DB:
		error, refreshToken := authorService.InsertRefreshToken(refreshToken)
		if error != nil {
			return exceptions.DatabaseConnectionException(error, c)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"result":        true,
			"user":          user,
			"access_token":  accessToken,
			"refresh_token": refreshToken.Token,
		})
	}

}

/**
 * Register User
 *
 * @param  user name
 * @param  password
 * @param  email
 * @param  image
 * @return  register user
 */
func Register(c echo.Context) error {
	var fuser forms.FUser
	fuser.Name = c.FormValue("username")
	fuser.Email = c.FormValue("email")
	fuser.Password = c.FormValue("password")
	storeErr, imageName := utils.SingleFileUpload(c)
	fuser.Image = imageName
	if storeErr != nil {
		return exceptions.StoreFileException(storeErr, c)
	}
	var validate = validator.New()
	err := validate.Struct(fuser)
	if err != nil {
		listError := exceptions.Validate(fuser)
		return exceptions.ValidationFieldException(listError, c)
	}

	fuser.Image = imageName
	// Check exits user in DB:
	_, user := authorService.GetUserByEmail(fuser.Email)
	if user.Name != "" {
		return exceptions.EmailExistsDBException(c)
	}
	// Register user
	var newUser entitys.User
	error, newUser := authorService.InsertUser(fuser.Name, fuser.Email, fuser.Password, fuser.Image)
	if error != nil {
		return exceptions.DatabaseConnectionException(error, c)
	}
	// Gen token to return for view:
	// Generate access_token
	accessToken := utils.GenerateJWT(user.Id, fuser.Name)
	// Generate refreshToken
	refreshToken := utils.GenerateRefreshToken(user.Id, fuser.Name)
	// Save refreshToken to DB:
	authorService.InsertRefreshToken(refreshToken)
	return c.JSON(http.StatusOK, echo.Map{
		"result":        true,
		"user":          newUser,
		"access_token":  accessToken,
		"refresh_token": refreshToken.Token,
	})

}
