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
		_, user = authorService.FindUserByUserEmail(fuser.Email)
		if (fuser.Email != user.Email) || (utils.CheckPasswordHash(fuser.Password, user.Password) != true) {
			return exceptions.IncorrectUserNamePasswordException(c)
		}
		// Generate access_token
		accessToken := utils.GenerateJWT(user.Id, user.Name)
		// Generate refreshToken
		refreshToken := utils.GenerateRefreshToken(user.Id, user.Name)
		// Save refreshToken to DB:
		error, refreshToken := authorService.SaveRefreshToken(refreshToken)
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
	_, user := authorService.FindUserByEmail(fuser.Email)
	if user.Name != "" {
		return exceptions.EmailExistsDBException(c)
	}
	// Register user
	var newUser entitys.User
	error, newUser := authorService.AddUser(fuser.Name, fuser.Email, fuser.Password, fuser.Image)
	if error != nil {
		return exceptions.DatabaseConnectionException(error, c)
	}
	// Gen token to return for view:
	// Generate access_token
	accessToken := utils.GenerateJWT(user.Id, fuser.Name)
	// Generate refreshToken
	refreshToken := utils.GenerateRefreshToken(user.Id, fuser.Name)
	// Save refreshToken to DB:
	authorService.SaveRefreshToken(refreshToken)
	return c.JSON(http.StatusOK, echo.Map{
		"result":        true,
		"user":          newUser,
		"access_token":  accessToken,
		"refresh_token": refreshToken.Token,
	})

}

func RenewToken(c echo.Context) error {
	var return_access_token string
	var return_refresh_token entitys.RefreshToken

	refreshToken := c.Param("refreshToken")
	//user_id := c.Param("user_id")
	//intVar, _ := strconv.Atoi(user_id)
	//_, user := authorService.FindUserById(intVar)

	if utils.ValidToken(refreshToken) {
		if utils.ExpiredToken(refreshToken) {
			_, refreshTokenObject := authorService.FindRefreshTokenByToken(refreshToken)
			if refreshTokenObject.UserName != "" {
				// Generate access_token
				return_access_token = utils.GenerateJWT(refreshTokenObject.UserId, refreshTokenObject.UserName)
				// Generate refreshToken
				return_refresh_token = utils.GenerateRefreshToken(refreshTokenObject.UserId, refreshTokenObject.UserName)
				// Save refreshToken to DB:
				authorService.SaveRefreshToken(return_refresh_token)
			}

		}
	}
	if return_access_token != "" && return_refresh_token.Token != "" {

		return c.JSON(http.StatusOK, echo.Map{
			"access_token":  return_access_token,
			"refresh_token": return_refresh_token.Token,
		})
	} else {
		return exceptions.InValidTokenException(c)
	}
}
