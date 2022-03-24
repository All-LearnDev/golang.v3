package userRepository

import (
	"projects/configs"
	"projects/entitys"
)

// save user and list image:
func AddUser(user entitys.User) entitys.User {
	Connection := configs.GetConnection()
	Connection.Create(&user)
	return user

}

func ListUser() []entitys.User {
	Connection := configs.GetConnection()
	var lists []entitys.User
	Connection.Find(&lists)
	Connection.Preload("Images").Find(&lists)
	return lists
}

func FindByUserId(Id int) entitys.User {
	Connection := configs.GetConnection()
	var user entitys.User
	Connection.First(&user, Id)
	Connection.Preload("Images").First(&user)
	return user
}
