package developerRepositoty

import (
	"projects/configs"
	"projects/entitys"
	"projects/repositorys/projectRepository"
)

func AddNewDeveloper(dev entitys.Developer) entitys.Developer {
	Connection := configs.GetConnection()
	dev = entitys.Developer{Name: dev.Name, Age: dev.Age, Projects: dev.Projects}
	Connection.Create(&dev)
	return dev
}

func GetDeveloperById(Id int) entitys.Developer {
	Connection := configs.GetConnection()
	var dev entitys.Developer
	Connection.Where("Id = ?", Id).Preload("Projects").First(&dev)
	//Connection.Preload("Images").First(&user)
	return dev
}

func ListDevelopers() []entitys.Developer {
	Connection := configs.GetConnection()
	var list []entitys.Developer
	Connection.Preload("Projects").Find(&list)
	return list
}

func DelDeveloperById(Id int) {
	Connection := configs.GetConnection()
	var developer entitys.Developer
	developer = GetDeveloperById(Id)
	// Remote relationship
	Connection.Model(&developer).Association("Projects").Clear()
	Connection.Delete(&entitys.Developer{}, Id)
}

func UpdateDeveloper(newDev entitys.Developer) entitys.Developer {
	Connection := configs.GetConnection()
	var developer entitys.Developer
	developer = GetDeveloperById(newDev.Id)
	developer.Name = newDev.Name
	developer.Age = newDev.Age
	var list []entitys.Project
	if developer.Name != "" {
		for i := 0; i < len(newDev.Projects); i++ {
			var project entitys.Project
			project = newDev.Projects[i]
			// Update project if exits or add new project
			if project.ID > 0 {
				project = projectRepository.UpdateProject(project.ID, project.Name, project.Customer)
			} else {
				project = projectRepository.SaveProject(project.Name, project.Customer)
			}
			list = append(list, project)
		}
		developer.Projects = list
		Connection.Save(&developer)
	}

	return developer
}
