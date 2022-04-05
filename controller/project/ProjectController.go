package project

import (
	"net/http"
	"projects/configs"
	"projects/entitys"
	"projects/exceptions"
	"projects/forms"
	"projects/repositorys/projectRepository"
	"projects/services/projectService"
	"projects/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func AddNewProject(c echo.Context) error {
	var fproject forms.FProject
	c.Bind(&fproject)
	var validate = validator.New()
	err := validate.Struct(fproject)
	if err != nil {
		listError := utils.Validate(fproject)
		return exceptions.ValidationFieldException(listError, c)
	} else {
		error, project := projectService.AddNewProject(fproject, c)
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

func ListProjects(c echo.Context) error {
	pg := configs.PaginateConfig()
	result := projectService.ListProjects()
	return c.JSON(200, pg.Response(result, c.Request(), &[]entitys.Project{}))
}

func FindProjectById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	list := projectRepository.FindProjectById(intVar)
	return c.JSON(http.StatusOK, list)
}

func FindSimpleProjectById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	list := projectRepository.FindSimpleProjectById(intVar)
	return c.JSON(http.StatusOK, list)
}
