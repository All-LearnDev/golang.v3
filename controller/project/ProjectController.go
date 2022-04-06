package project

import (
	"net/http"
	"projects/configs"
	"projects/entitys"
	"projects/exceptions"
	"projects/forms"
	"projects/services/projectService"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

/**
* Insert new project to Database.
*
* @param  project'name
* @param  project'description
* @return  project or error when connection to database fail
 */
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

/**
* Update project info.
*
* @param  project'id
* @param  project'name
* @param  project'description
* @return  project or error when connection to database fail
 */
func UpdateProject(c echo.Context) error {
	var fproject forms.FProject
	c.Bind(&fproject)
	var validate = validator.New()
	err := validate.Struct(fproject)
	if err != nil {
		listError := exceptions.Validate(fproject)
		return exceptions.ValidationFieldException(listError, c)
	} else {
		error, project := projectService.UpdateProject(fproject, c)
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

/**
* Get list projects from database.
* @return  list projects or error when connection to database fail
 */
func GetListProjects(c echo.Context) error {
	pg := configs.PaginateConfig()
	result := projectService.GetListProjects()
	return c.JSON(200, pg.Response(result, c.Request(), &[]entitys.Project{}))
}

/**
* Get project info.
*
* @param  project'id
* @return  project or error when connection to database fail
 */
func GetProjectById(c echo.Context) error {
	id := c.Param("id")
	id_project, _ := strconv.Atoi(id)
	error, project := projectService.GetProjectById(id_project)
	if error != nil {
		return exceptions.DatabaseConnectionException(error, c)
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"result":  true,
			"project": project,
		})
	}
}

/**
* Delete project.
*
* @param  project'id
* @return  nil or error when connection to database fail
 */
func DeleteProjectById(c echo.Context) error {
	id := c.Param("id")
	id_project, _ := strconv.Atoi(id)
	error, _ := projectService.GetProjectById(id_project)
	if error != nil {
		return exceptions.RecordNotFoundException(error, c)
	} else {
		error := projectService.DeleteProjectById(id_project)
		if error != nil {
			return exceptions.DatabaseConnectionException(error, c)
		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"result": true,
			})
		}
	}
}
