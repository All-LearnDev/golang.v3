package authorService

import (
	"projects/entitys"
	"projects/forms"
	"projects/repositorys/authorRepository"
)

func AddJUser(name string, email string, password string) entitys.JUser {
	return authorRepository.AddJUser(name, email, password)

}

func AddRolesToUser(user entitys.JUser, fRoles []forms.Role) (error, entitys.JUser) {
	return authorRepository.AddRolesToUser(user, fRoles)
}

func FindUserByEmail(email string) (error, entitys.JUser) {
	return authorRepository.FindUserByEmail(email)
}

func FindUserById(id int) (error, entitys.JUser) {
	return authorRepository.FindUserById(id)

}

func FindUserByUserName(name string) (error, entitys.JUser) {
	return authorRepository.FindUserByUserName(name)
}

func FindRefreshTokenByUserId(Id int) (error, entitys.RefreshToken) {
	return authorRepository.FindRefreshTokenByUserId(Id)

}

func FindRefreshTokenByUserName(userName string) entitys.RefreshToken {

	return authorRepository.FindRefreshTokenByUserName(userName)

}

func FindRefreshTokenByToken(token string) (error, entitys.RefreshToken) {
	return authorRepository.FindRefreshTokenByToken(token)

}

func SaveRefreshToken(refreshToken entitys.RefreshToken) (error, entitys.RefreshToken) {
	return authorRepository.SaveRefreshToken(refreshToken)

}
