package repositorys

import (
	"projects/configs"
	"projects/entitys"
	"projects/utils"
)

func AddJUser(name string, email string, password string) entitys.JUser {
	Connection := configs.GetConnection()
	var hash string
	hash, _ = utils.HashPassword(password)
	user := entitys.JUser{Name: name, Email: email, Password: hash}
	Connection.Create(&user)
	return user

}

func FindUserByEmail(email string) entitys.JUser {
	Connection := configs.GetConnection()
	var user entitys.JUser
	Connection.Where("Email = ?", email).First(&user)
	return user
}

func FindUserByUserName(name string) entitys.JUser {
	Connection := configs.GetConnection()
	var user entitys.JUser
	Connection.Where("Name = ?", name).First(&user)
	return user
}

func FindRefreshTokenByUserId(Id int) entitys.RefreshToken {
	Connection := configs.GetConnection()
	var token entitys.RefreshToken
	Connection.Where("Id = ?", Id).First(&token)
	return token
}

func FindRefreshTokenByUserName(userName string) entitys.RefreshToken {
	Connection := configs.GetConnection()
	var token entitys.RefreshToken
	Connection.First(&token, "UserName = ?", userName)
	return token
}

func FindRefreshTokenByToken(token string) entitys.RefreshToken {
	Connection := configs.GetConnection()
	var refreshToken entitys.RefreshToken
	Connection.Where("Token = ?", token).First(&refreshToken)
	return refreshToken
}

func SaveRefreshToken(refreshToken entitys.RefreshToken) entitys.RefreshToken {
	Connection := configs.GetConnection()
	Connection.Create(&refreshToken)
	return refreshToken

}
