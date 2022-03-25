package developer

import (
	"net/http"
	"projects/entitys"
	"projects/forms"
	"projects/services/developerService"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddNewDeveloper(c echo.Context) error {
	var fdeveloper forms.FDeveloper
	c.Bind(&fdeveloper)
	var developer entitys.Developer
	developer.Name = fdeveloper.Name
	developer.Age = fdeveloper.Age
	var list []entitys.Project
	if len(fdeveloper.List) > 0 {
		for i := 0; i < len(fdeveloper.List); i++ {
			project := entitys.Project{Name: fdeveloper.List[i].Name, Customer: fdeveloper.List[i].Customer}
			list = append(list, project)
		}
		developer.Projects = list
	}
	developer = developerService.AddNewDeveloper(developer)
	return c.JSON(http.StatusOK, developer)
}

func GetDeveloperById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	dev := developerService.GetDeveloperById(intVar)
	return c.JSON(http.StatusOK, dev)
}

func ListDevelopers(c echo.Context) error {
	//  []entitys.Developer
	list := developerService.ListDevelopers()
	return c.JSON(http.StatusOK, list)
}

func DelDeveloperById(c echo.Context) error {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)
	developerService.DelDeveloperById(intVar)
	return c.String(http.StatusOK, "Deleted !")

}

func UpdateDeveloper(c echo.Context) error {
	var fdeveloper forms.FDeveloper
	c.Bind(&fdeveloper)
	var developer entitys.Developer
	developer.Name = fdeveloper.Name
	developer.Age = fdeveloper.Age
	developer.Id = fdeveloper.Id
	var list []entitys.Project
	if len(fdeveloper.List) > 0 {
		for i := 0; i < len(fdeveloper.List); i++ {
			project := entitys.Project{ID: fdeveloper.List[i].ID, Name: fdeveloper.List[i].Name, Customer: fdeveloper.List[i].Customer}
			list = append(list, project)
		}
		developer.Projects = list
	}
	developer = developerService.UpdateDeveloper(developer)
	return c.JSON(http.StatusOK, developer)

}
