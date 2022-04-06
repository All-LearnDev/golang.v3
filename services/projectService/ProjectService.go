package projectService

import (
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/projectRepository"
	"projects/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InsertNewProject(project forms.FProject, c echo.Context) (error, entitys.Project) {

	auth := c.Request().Header.Get("Authorization")
	jwt := strings.Split(auth, " ")[1]
	user := utils.GetUserFromTokden(jwt)
	return projectRepository.InsertNewProject(user, project)
}

func GetListProjects() (tx *gorm.DB) {
	return projectRepository.GetListProjects()
}

func GetProjectById(Id int) entitys.Project {
	return projectRepository.GetProjectById(Id)
}

func GetSimpleProjectById(Id int) entitys.Project {
	return projectRepository.GetSimpleProjectById(Id)
}
