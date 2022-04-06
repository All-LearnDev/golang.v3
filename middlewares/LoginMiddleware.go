package middlewares

import (
	"projects/exceptions"
	"projects/repositorys/authorRepository"
	"projects/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

/**
* Middleware to allow user access or not.
*
* @param  jwt token from client
* @return allow user access resource/ api or not
 */
func LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if len(auth) < 10 {
			return exceptions.InValidTokenException(c)
		}
		if len(auth) > 10 {
			jwt := strings.Split(auth, " ")[1]
			if utils.IsValidToken(jwt) == false {
				return exceptions.InValidTokenException(c)
			}
			user := utils.GetUserFromTokden(jwt)
			_, dbuser := authorRepository.GetUserById(user.Id)
			//println(" DB id ", dbuser.Id, " DB name ", dbuser.Name)
			//println(" Token user  id ", dbuser.Id, " Token user  name ", dbuser.Name)
			if (dbuser.Id != user.Id) || (dbuser.Name != user.Name) {

				return exceptions.UnauthorizedException(c)
			}
		}
		return next(c)
	}
}
