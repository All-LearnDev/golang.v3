package developerService

import (
	"projects/entitys"
	"projects/repositorys/developerRepositoty"
)

func AddNewDeveloper(dev entitys.Developer) entitys.Developer {
	return developerRepositoty.AddNewDeveloper(dev)
}

func GetDeveloperById(Id int) entitys.Developer {
	return developerRepositoty.GetDeveloperById(Id)
}

func ListDevelopers() []entitys.Developer {
	return developerRepositoty.ListDevelopers()
}
