package main

import (
	"fmt"
	"net/http"
	"projects/configs"
	"projects/routers"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(rest *echo.Echo, err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

}

func main() {

	configs.AutoMigrate()
	rest := echo.New()
	routers.InitializeApiMapping(rest)
	rest.Logger.Fatal(rest.Start(":1323"))

}
