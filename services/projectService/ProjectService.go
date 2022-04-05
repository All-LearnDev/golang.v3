package projectService

import (
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/projectRepository"

	"gorm.io/gorm"
)

func FindProjectById(Id int) entitys.Project {
	return projectRepository.FindProjectById(Id)
}

func FindSimpleProjectById(Id int) entitys.Project {
	return projectRepository.FindSimpleProjectById(Id)
}

func ListProjects() (tx *gorm.DB) {
	return projectRepository.ListProjects()
}

func AddNewProject(project forms.FProject) (error, entitys.Project) {
	return projectRepository.AddNewProject(project)
}

func ListEagerProjects() []entitys.Project {
	return projectRepository.ListEagerProjects()
}
