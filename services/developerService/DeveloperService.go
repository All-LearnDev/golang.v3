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
func DelDeveloperById(Id int) {
	developerRepositoty.DelDeveloperById(Id)
}

func UpdateDeveloper(dev entitys.Developer) entitys.Developer {

	return developerRepositoty.UpdateDeveloper(dev)
}
