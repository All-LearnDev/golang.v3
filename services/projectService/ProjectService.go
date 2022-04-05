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

func AddNewProject(project forms.FProject, c echo.Context) (error, entitys.Project) {

	auth := c.Request().Header.Get("Authorization")
	jwt := strings.Split(auth, " ")[1]
	user := utils.GetUserFromTokden(jwt)
	return projectRepository.AddNewProject(user, project)
}

func ListProjects() (tx *gorm.DB) {
	return projectRepository.ListProjects()
}

func FindProjectById(Id int) entitys.Project {
	return projectRepository.FindProjectById(Id)
}

func FindSimpleProjectById(Id int) entitys.Project {
	return projectRepository.FindSimpleProjectById(Id)
}
