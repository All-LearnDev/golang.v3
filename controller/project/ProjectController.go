package project

import (
	"net/http"
	"projects/configs"
	"projects/entitys"
	"projects/exceptions"
	"projects/forms"
	"projects/repositorys/projectRepository"
	"projects/services/projectService"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func InsertNewProject(c echo.Context) error {
	var fproject forms.FProject
	c.Bind(&fproject)
	var validate = validator.New()
	err := validate.Struct(fproject)
	if err != nil {
		listError := exceptions.Validate(fproject)
		return exceptions.ValidationFieldException(listError, c)
	} else {
		error, project := projectService.InsertNewProject(fproject, c)
		if error != nil {
			return exceptions.DatabaseConnectionException(error, c)
		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"result":  true,
				"project": project,
			})
		}
	}
}

func GetListProjects(c echo.Context) error {
	pg := configs.PaginateConfig()
	result := projectService.GetListProjects()
	return c.JSON(200, pg.Response(result, c.Request(), &[]entitys.Project{}))
}

func GetProjectById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	list := projectRepository.GetProjectById(intVar)
	return c.JSON(http.StatusOK, list)
}

func GetSimpleProjectById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	list := projectService.GetSimpleProjectById(intVar)
	return c.JSON(http.StatusOK, list)
}
