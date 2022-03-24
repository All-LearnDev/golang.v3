package main

import (
	"projects/configs"
	"projects/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.AutoMigrate()
	rest := echo.New()
	routers.InitializeApiMapping(rest)
	rest.Logger.Fatal(rest.Start(":1323"))

}
