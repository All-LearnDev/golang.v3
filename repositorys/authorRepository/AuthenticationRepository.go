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
	user := entitys.JUser{Name: name, Email: email, Password: hash}
	Connection.Create(&user)
	return user

}

func AddRolesToUser(user entitys.JUser, fRoles []forms.Role) entitys.JUser {
	var roles []entitys.Roles
	mapper := dto.Mapper{}
	mapper.Map(&roles, fRoles)
	user.Roles = roles
	Connection.Save(user)
	return user

}

func FindUserByEmail(email string) entitys.JUser {

	var user entitys.JUser
	Connection.Where("Email = ?", email).Preload("Roles").First(&user)
	return user
}

func FindUserById(id int) entitys.JUser {

	var user entitys.JUser
	Connection.Where("Id = ?", id).Preload("Roles").First(&user)
	return user
}

func FindUserByUserName(name string) entitys.JUser {
	var user entitys.JUser
	Connection.Where("Name = ?", name).Preload("Roles").First(&user)
	return user
}

func FindRefreshTokenByUserId(Id int) entitys.RefreshToken {
	var token entitys.RefreshToken
	Connection.Where("Id = ?", Id).First(&token)
	return token
}

func FindRefreshTokenByUserName(userName string) entitys.RefreshToken {
	var token entitys.RefreshToken
	Connection.First(&token, "UserName = ?", userName)
	return token
}

func FindRefreshTokenByToken(token string) entitys.RefreshToken {
	var refreshToken entitys.RefreshToken
	Connection.Where("Token = ?", token).First(&refreshToken)
	return refreshToken
}

func SaveRefreshToken(refreshToken entitys.RefreshToken) entitys.RefreshToken {
	Connection.Create(&refreshToken)
	return refreshToken

}
