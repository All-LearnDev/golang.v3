package projectRepository

import (
	"projects/configs"
	"projects/entitys"
	"projects/forms"

	"gorm.io/gorm"
)

var Connection = configs.GetConnection()

func AddNewProject(user entitys.User, project forms.FProject) (error, entitys.Project) {
	newProject := entitys.Project{
		Name:           project.Name,
		Description:    project.Description,
		CreateByUserId: user.Id}
	error := Connection.Save(&newProject).Error
	return error, newProject
}

func ListProjects() (tx *gorm.DB) {
	var list []entitys.Project
	result := Connection.Find(&list)
	Connection.Find(&list)
	return result
}

func FindProjectById(Id int) entitys.Project {

	var project entitys.Project
	Connection.Where("Id = ?", Id).Preload("Projects").First(&project)
	return project

}

func FindSimpleProjectById(Id int) entitys.Project {

	var project entitys.Project
	Connection.Where("Id = ?", Id).First(&project)
	return project

}

func UpdateProject(id int, name string, customer string) entitys.Project {
	project := FindSimpleProjectById(id)
	project.Name = name
	Connection.Save(&project)
	return project
}
