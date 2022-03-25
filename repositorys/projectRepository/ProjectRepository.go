package projectRepository

import (
	"projects/configs"
	"projects/entitys"
)

func FindProjectById(Id int) entitys.Project {
	Connection := configs.GetConnection()
	var project entitys.Project
	Connection.Where("Id = ?", Id).Preload("Projects").First(&project)
	//Connection.Preload("Images").First(&user)
	return project

}

func FindSimpleProjectById(Id int) entitys.Project {
	Connection := configs.GetConnection()
	var project entitys.Project
	Connection.Where("Id = ?", Id).First(&project)
	//Connection.Preload("Images").First(&user)
	return project

}

func ListProjects() []entitys.Project {
	Connection := configs.GetConnection()
	var list []entitys.Project
	Connection.Find(&list)
	return list
}

func ListEagerProjects() []entitys.Project {
	Connection := configs.GetConnection()
	var list []entitys.Project
	Connection.Preload("Developers").Find(&list)
	return list
}

func DelProjectById(Id int) {
	Connection := configs.GetConnection()
	var project entitys.Project
	project = FindProjectById(Id)
	// Remote relationship
	Connection.Model(&project).Association("Developers").Clear()
	Connection.Delete(&entitys.Project{}, Id)
}

func UpdateProject(id int, name string, customer string) entitys.Project {
	project := FindSimpleProjectById(id)
	project.Name = name
	project.Customer = customer
	Connection := configs.GetConnection()
	Connection.Save(&project)
	return project
}

func SaveProject(name string, customer string) entitys.Project {
	project := entitys.Project{Name: name, Customer: customer}
	Connection := configs.GetConnection()
	Connection.Save(&project)
	return project
}
