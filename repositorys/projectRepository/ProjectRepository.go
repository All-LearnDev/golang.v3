package projectRepository

import (
	"projects/configs"
	"projects/entitys"
)

var Connection = configs.GetConnection()

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

func ListProjects() []entitys.Project {
	var list []entitys.Project
	Connection.Find(&list)
	return list
}

func ListEagerProjects() []entitys.Project {
	var list []entitys.Project
	Connection.Preload("Developers").Find(&list)
	return list
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
	project.Customer = customer
	Connection.Save(&project)
	return project
}

func SaveProject(name string, customer string) entitys.Project {
	project := entitys.Project{Name: name, Customer: customer}
	Connection.Save(&project)
	return project
}
