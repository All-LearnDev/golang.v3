package utils

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Process is the middleware function.
func AdminProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token !")
		}
		jwt := strings.Split(auth, " ")[1]
		if ValidToken(jwt) == false {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token !")
		}
		user := GetRolesFromToken(jwt)
		//println(user.Roles[0].Name)
		//println(user.Roles[1].Name)
		if len(user.Roles) > 0 {
			if user.Roles[0].Name != "Admin" {
				return echo.NewHTTPError(http.StatusForbidden, "Anauthorized")
			}
		}

		return next(c)
	}
}

func ViewerProcess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token !")
		}
		jwt := strings.Split(auth, " ")[1]
		if ValidToken(jwt) == false {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token !")
		}
		user := GetRolesFromToken(jwt)
		//println(user.Roles[0].Name)
		//	println(user.Roles[1].Name)
		if len(user.Roles) > 0 {
			if user.Roles[0].Name != "Viewer" {
				return echo.NewHTTPError(http.StatusForbidden, "Anauthorized")
			}
		}

		return next(c)
	}
}
