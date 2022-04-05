package routers

import (
	authorController "projects/controller/author"
	"projects/controller/project"
	"projects/middlewares"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func InitializeApiMapping(rest *echo.Echo) {
	// Enable cors in Echo:
	rest.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	authGroup := rest.Group("/author")
	authGroup.POST("/register", authorController.Register)
	authGroup.POST("/login", authorController.Login)
	authGroup.GET("/renew/:refreshToken", authorController.RenewToken)

	projectGroup := rest.Group("/project", middlewares.LoginMiddleware)
	projectGroup.GET("/list", project.ListProjects)
	projectGroup.POST("/add", project.AddNewProject)
	projectGroup.GET("/lazy/findbyid/:id", project.FindSimpleProjectById)
	projectGroup.GET("/eager/findbyid/:id", project.FindProjectById)

}
