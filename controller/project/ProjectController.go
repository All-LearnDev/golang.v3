package project

import (
	"net/http"
	"projects/repositorys/projectRepository"
	"projects/services/projectService"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

func ListProjects(c echo.Context) error {
	list := projectService.ListProjects()
	return c.JSON(http.StatusOK, list)

}

func ListEagerProjects(c echo.Context) error {
	list := projectService.ListEagerProjects()
	return c.JSON(http.StatusOK, list)

}
