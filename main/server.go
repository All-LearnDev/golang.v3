package main

import (
	"fmt"
	"os"
	"projects/configs"
	"projects/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("editor:", os.Getenv("EDITOR"))
	configs.AutoMigrate()
	rest := echo.New()
	rest.Use(middleware.CORS())
	routers.InitializeApiMapping(rest)
	rest.Logger.Fatal(rest.Start(":1323"))

}
