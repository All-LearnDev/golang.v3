package projectRepository

import (
	"projects/configs"
	"projects/entitys"
	"projects/forms"

	"gorm.io/gorm"
)

var Connection = configs.GetConnection()

/**
* Insert new project to Database.
*
* @param  project'name
* @param  project'description
* @return  project or error when connection to database fail
 */
func InsertNewProject(user entitys.User, project forms.FProject) (error, entitys.Project) {
	newProject := entitys.Project{
		Name:           project.Name,
		Description:    project.Description,
		CreateByUserId: user.Id}
	error := Connection.Save(&newProject).Error
	return error, newProject
}

/**
* Get list projects from database.
* @return  list projects or error when connection to database fail
 */
func GetListProjects() (tx *gorm.DB) {
	var list []entitys.Project
	result := Connection.Find(&list)
	Connection.Find(&list)
	return result
}

/**
* Update project info.
*
* @param  user'id
* @param  project'info
* @return  project or error when connection to database fail
 */
func UpdateProject(project entitys.Project) (error, entitys.Project) {
	error := Connection.Save(&project).Error
	return error, project
}

/**
* Get project info.
*
* @param  project'id
* @return  project or error when connection to database fail
 */
func GetProjectById(Id int) (error, entitys.Project) {
	var project entitys.Project
	error := Connection.Where("Id = ?", Id).First(&project).Error
	return error, project
}

/**
* Get project info.
*
* @param  project'id
* @return  project or error when connection to database fail
 */
func DeleteProjectById(Id int) error {
	error := Connection.Delete(&entitys.Project{}, Id).Error
	return error
}
