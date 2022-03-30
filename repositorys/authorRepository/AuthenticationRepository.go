package authorRepository

import (
	"projects/configs"
	"projects/entitys"
	"projects/forms"
	"projects/utils"

	"github.com/dranikpg/dto-mapper"
)

var Connection = configs.GetConnection()

func AddJUser(name string, email string, password string) entitys.JUser {
	var hash string
	hash, _ = utils.HashPassword(password)
	print(" name", name)
	user := entitys.JUser{Name: name, Email: email, Password: hash}
	Connection.Create(&user)
	return user

}

func AddRolesToUser(user entitys.JUser, fRoles []forms.Role) (error, entitys.JUser) {
	var roles []entitys.Roles
	mapper := dto.Mapper{}
	mapper.Map(&roles, fRoles)
	user.Roles = roles
	error := Connection.Save(user).Error
	return error, user

}

func FindUserByEmail(email string) (error, entitys.JUser) {

	var user entitys.JUser
	error := Connection.Where("Email = ?", email).Preload("Roles").First(&user).Error
	return error, user
}

func FindUserById(id int) (error, entitys.JUser) {

	var user entitys.JUser
	error := Connection.Where("Id = ?", id).Preload("Roles").First(&user).Error
	return error, user
}

func FindUserByUserName(name string) (error, entitys.JUser) {
	var user entitys.JUser
	error := Connection.Where("Name = ?", name).Preload("Roles").First(&user).Error
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
