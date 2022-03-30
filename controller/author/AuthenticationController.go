package controller

import (
	"net/http"
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/authorRepository"
	"strconv"

	"projects/utils"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	var user entitys.JUser
	user = authorRepository.FindUserByUserName(username)
	if (username != user.Name) || (utils.CheckPasswordHash(password, user.Password) != true) {
		return echo.ErrUnauthorized
	}
	// Generate access_token
	er, accessToken := utils.GenerateJWT(username, user.Roles)
	// Generate refreshToken
	er, refreshToken := utils.GenerateRefreshToken(username, user.Roles)
	// Save refreshToken to DB:
	authorRepository.SaveRefreshToken(refreshToken)
	if er != nil {
		return er
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"user":          user,
			"access_token":  accessToken,
			"refresh_token": refreshToken.Token,
		})
	}

}

func AddRolesToUser(c echo.Context) error {

	var userRole forms.FUserRole
	c.Bind(&userRole)
	user_id := userRole.UserId
	// Finf user by id:
	user := authorRepository.FindUserById(user_id)
	// Save roles to DB:
	user = authorRepository.AddRolesToUser(user, userRole.Roles)
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	// Check exits user in DB:
	user := authorRepository.FindUserByEmail(email)
	if user.Name != "" {
		return c.String(http.StatusBadRequest, "Email exits in DB")
	}
	// Register user
	var newUser entitys.JUser
	newUser = authorRepository.AddJUser(username, email, password)
	// Gen token to return for view:
	_, accessToken := utils.GenerateJWT(username, user.Roles)
	// Generate access_token
	// Generate refreshToken
	_, refreshToken := utils.GenerateRefreshToken(username, user.Roles)
	// Save refreshToken to DB:
	authorRepository.SaveRefreshToken(refreshToken)
	return c.JSON(http.StatusOK, echo.Map{
		"user":          newUser,
		"access_token":  accessToken,
		"refresh_token": refreshToken.Token,
	})

}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"token": nil,
	})
}

func RenewToken(c echo.Context) error {
	var return_access_token string
	var return_refresh_token entitys.RefreshToken

	refreshToken := c.Param("refreshToken")
	user_id := c.Param("user_id")
	intVar, _ := strconv.Atoi(user_id)
	user := authorRepository.FindUserById(intVar)

	if utils.ValidToken(refreshToken) {
		if utils.ExpiredToken(refreshToken) {
			refreshTokenObject := authorRepository.FindRefreshTokenByToken(refreshToken)
			if refreshTokenObject.UserName != "" {
				// Generate access_token
				_, return_access_token = utils.GenerateJWT(refreshTokenObject.UserName, user.Roles)
				// Generate refreshToken
				_, return_refresh_token = utils.GenerateRefreshToken(refreshTokenObject.UserName, user.Roles)
				// Save refreshToken to DB:
				authorRepository.SaveRefreshToken(return_refresh_token)
			}

		}
	}
	if return_access_token != "" && return_refresh_token.Token != "" {

		return c.JSON(http.StatusOK, echo.Map{
			"access_token":  return_access_token,
			"refresh_token": return_refresh_token.Token,
		})
	} else {
		return c.String(http.StatusBadRequest, " Invalid token ! ")
	}
}

func ExpireToken(c echo.Context) error {
	access_token := c.Param("access_token")
	flag := utils.ExpiredToken(access_token)
	return c.JSON(http.StatusOK, echo.Map{
		"token":        access_token,
		"expire_token": flag,
	})
}

func ValidToken(c echo.Context) error {
	access_token := c.Param("access_token")
	result := utils.ValidToken(access_token)
	return c.JSON(http.StatusOK, echo.Map{
		"token":        access_token,
		"expire_token": result,
	})
}
