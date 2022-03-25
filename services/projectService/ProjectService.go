package projectService

import (
	"projects/entitys"
	"projects/repositorys/projectRepository"
)

func FindProjectById(Id int) entitys.Project {
	return projectRepository.FindProjectById(Id)
}

func FindSimpleProjectById(Id int) entitys.Project {
	return projectRepository.FindSimpleProjectById(Id)
}

func ListProjects() []entitys.Project {
	return projectRepository.ListProjects()
}

func ListEagerProjects() []entitys.Project {
	return projectRepository.ListEagerProjects()
}
