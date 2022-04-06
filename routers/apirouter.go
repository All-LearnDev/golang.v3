package routers

import (
	authorController "projects/controller/author"
	"projects/controller/project"
	"projects/middlewares"

	echo "github.com/labstack/echo/v4"
)

func InitializeApiMapping(rest *echo.Echo) {

	authGroup := rest.Group("/author")
	authGroup.POST("/register", authorController.Register)
	authGroup.POST("/login", authorController.Login)

	projectGroup := rest.Group("/project", middlewares.LoginMiddleware)
	projectGroup.POST("/add", project.InsertNewProject)
	projectGroup.POST("/update", project.UpdateProject)
	projectGroup.GET("/getbyid/:id", project.GetProjectById)
	projectGroup.GET("/delete/:id", project.DeleteProjectById)
	// list/paging?size=3&page=0&sort=-name
	projectGroup.GET("/list/paging", project.GetListProjects)

}
