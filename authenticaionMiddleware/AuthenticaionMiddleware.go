package middleware

import (
	"net/http"
	"projects/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

// Process is the middleware function.
func LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"token_valid": false,
				"message":     "Invalid token ",
			})

		}
		jwt := strings.Split(auth, " ")[1]
		if utils.ValidToken(jwt) == false {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"token_valid": false,
				"message":     "Invalid token ",
			})
		}
		user := utils.GetUserNameFromToken(jwt)
		if len(user.Name) == 0 {
			return c.JSON(http.StatusForbidden, echo.Map{
				"token_valid": true,
				"message":     "Anauthorized ",
			})
		}
		return next(c)
	}
}
