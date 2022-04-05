package middlewares

import (
	"projects/exceptions"
	"projects/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if len(auth) < 10 {
			return exceptions.InValidTokenException(c)
		}
		if len(auth) > 10 {
			jwt := strings.Split(auth, " ")[1]
			if utils.ValidToken(jwt) == false {
				return exceptions.InValidTokenException(c)
			}
			user := utils.GetUserNameFromToken(jwt)
			if len(user.Name) == 0 {
				return exceptions.UnauthorizedException(c)
			}
		}
		return next(c)
	}
}
