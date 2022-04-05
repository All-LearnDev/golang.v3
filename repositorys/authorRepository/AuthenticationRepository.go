package authorRepository

import (
	"projects/configs"
	"projects/entitys"
	"projects/utils"
)

var Connection = configs.GetConnection()

func AddUser(name string, email string, password string, image string) (error, entitys.User) {
	var hash string
	hash, _ = utils.HashPassword(password)
	user := entitys.User{Name: name, Email: email, Password: hash, Image: image}
	error := Connection.Create(&user).Error
	return error, user

}

func FindUserByEmail(email string) (error, entitys.User) {

	var user entitys.User
	error := Connection.Where("Email = ?", email).First(&user).Error
	return error, user
}

func FindUserById(id int) (error, entitys.User) {

	var user entitys.User
	error := Connection.Where("Id = ?", id).First(&user).Error
	return error, user
}

func FindUserByUserName(name string) (error, entitys.User) {
	var user entitys.User
	error := Connection.Where("Name = ?", name).First(&user).Error
	return error, user
}

func FindUserByUserEmail(email string) (error, entitys.User) {
	var user entitys.User
	error := Connection.Where("Email = ?", email).First(&user).Error
	return error, user
}

func FindRefreshTokenByUserId(Id int) (error, entitys.RefreshToken) {
	var token entitys.RefreshToken
	error := Connection.Where("Id = ?", Id).First(&token).Error
	return error, token
}

func FindRefreshTokenByUserName(userName string) entitys.RefreshToken {
	var token entitys.RefreshToken
	Connection.First(&token, "UserName = ?", userName)
	return token
}

func FindRefreshTokenByToken(token string) (error, entitys.RefreshToken) {
	var refreshToken entitys.RefreshToken
	error := Connection.Where("Token = ?", token).First(&refreshToken).Error
	return error, refreshToken
}

func SaveRefreshToken(refreshToken entitys.RefreshToken) (error, entitys.RefreshToken) {
	error := Connection.Create(&refreshToken).Error
	return error, refreshToken

}
