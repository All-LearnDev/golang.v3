package developer

import (
	"net/http"
	"projects/entitys"
	developerService "projects/services/developerservice"

	"github.com/labstack/echo/v4"
)

func AddNewDeveloper(c echo.Context) error {
	var developer entitys.Developer
	c.Bind(&developer)
	developer = developerService.AddNewDeveloper(developer)
	return c.JSON(http.StatusOK, developer)
}

func GetDeveloperById(c echo.Context) error {
	return c.JSON(http.StatusOK, " kon meo ! ")
}

func ListDevelopers(c echo.Context) error {
	//  []entitys.Developer
	list := developerService.ListDevelopers()
	return c.JSON(http.StatusOK, list)
}
