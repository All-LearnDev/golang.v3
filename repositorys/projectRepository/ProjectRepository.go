package projectRepository

import (
	"projects/configs"
	"projects/entitys"
	"projects/forms"

	"gorm.io/gorm"
)

var Connection = configs.GetConnection()

func AddNewProject(project forms.FProject) (error, entitys.Project) {
	newProject := entitys.Project{Name: project.Name, Description: project.Description}
	error := Connection.Save(&newProject).Error
	return error, newProject
}

func ListProjects() (tx *gorm.DB) {
	var list []entitys.Project
	result := Connection.Find(&list)
	Connection.Find(&list)
	return result
}

func ListEagerProjects() []entitys.Project {
	var list []entitys.Project
	Connection.Preload("Developers").Find(&list)
	return list
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

func DelProjectById(Id int) {
	var project entitys.Project
	project = FindProjectById(Id)
	Connection.Model(&project).Association("Developers").Clear()
	Connection.Delete(&entitys.Project{}, Id)
}

func UpdateProject(id int, name string, customer string) entitys.Project {
	project := FindSimpleProjectById(id)
	project.Name = name
	Connection.Save(&project)
	return project
}

func SaveProject(name string, customer string) entitys.Project {
	project := entitys.Project{Name: name, Description: customer}
	Connection.Save(&project)
	return project
}
