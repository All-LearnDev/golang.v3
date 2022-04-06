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

/**
* Insert new project to Database.
*
* @param  project'name
* @param  project'description
* @return  project or error when connection to database fail
 */
func InsertNewProject(project forms.FProject, c echo.Context) (error, entitys.Project) {
	auth := c.Request().Header.Get("Authorization")
	jwt := strings.Split(auth, " ")[1]
	user := utils.GetUserFromTokden(jwt)
	return projectRepository.InsertNewProject(user, project)
}

/**
* Insert new project to Database.
*
* @param  project'name
* @param  project'description
* @return  project or error when connection to database fail
 */
func UpdateProject(form forms.FProject, c echo.Context) (error, entitys.Project) {
	auth := c.Request().Header.Get("Authorization")
	jwt := strings.Split(auth, " ")[1]
	user := utils.GetUserFromTokden(jwt)
	error, project := projectRepository.GetProjectById(form.ID)
	println(" id project ", project.ID, " Name ", project.Name)
	if error == nil {
		project.ID = form.ID
		project.Name = form.Name
		project.Description = form.Description
		project.UpdateByUserId = user.Id
		_, project = projectRepository.UpdateProject(project)
		println(" id project ", project.ID, " Name ", project.Name)
		return error, project
	} else {
		return error, entitys.Project{}
	}

}

/**
* Get list projects from database.
*@return  list projects or error when connection to database fail
 */
func GetListProjects() (tx *gorm.DB) {
	return projectRepository.GetListProjects()
}

/**
* Get project info.
*
* @param  project'id
* @return  project or error when connection to database fail
 */
func GetProjectById(Id int) (error, entitys.Project) {
	return projectRepository.GetProjectById(Id)
}

/**
* Get project info.
*
* @param  project'id
* @return  nil or error when connection to database fail
 */
func DeleteProjectById(Id int) error {
	return projectRepository.DeleteProjectById(Id)
}
