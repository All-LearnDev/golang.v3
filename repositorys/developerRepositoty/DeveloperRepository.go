package developerRepositoty

import (
	"projects/configs"
	"projects/entitys"
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
