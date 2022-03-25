package userService

import (
	"projects/entitys"
	"projects/repositorys/userRepository"
)

func AddUser(entity entitys.User) entitys.User {
	user := userRepository.AddUser(entity)
	return user
}

func ListUser() []entitys.User {
	return userRepository.ListUser()
}

func ListLazyUser() []entitys.User {
	return userRepository.ListUser()
}

func FindByUserId(Id int) entitys.User {
	return userRepository.FindByUserId(Id)
}

func DeleteUserById(Id int) {
	userRepository.DeleteUserById(Id)
}
