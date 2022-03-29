package userRepository

import (
	"projects/configs"
	"projects/entitys"
)

var Connection = configs.GetConnection()

// save user and list image:
func AddUser(user entitys.User) entitys.User {

	Connection.Create(&user)
	return user

}

func ListUser() []entitys.User {
	var lists []entitys.User
	Connection.Find(&lists)
	Connection.Preload("Images").Find(&lists)
	return lists
}

func ListLazyUser() []entitys.User {
	var lists []entitys.User
	Connection.Find(&lists)
	return lists
}

func FindByUserId(Id int) entitys.User {
	var user entitys.User
	Connection.First(&user, Id)
	Connection.Preload("Images").First(&user)
	return user
}

func DeleteUserById(Id int) {
	var user entitys.User
	user = FindByUserId(Id)
	Connection.Select("Images").Delete(&user)
	Connection.Delete(&user)
}
