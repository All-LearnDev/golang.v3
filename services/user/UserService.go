package services

import (
	"projects/entitys"
	repositorys "projects/repositorys/user"
)

func AddUser(entity entitys.User) entitys.User {
	user := repositorys.AddUser(entity)
	return user
}

func ListUser() []entitys.User {
	return repositorys.ListUser()
}

func FindByUserId(Id int) entitys.User {
	return repositorys.FindByUserId(Id)
}
