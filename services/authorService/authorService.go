package authorService

import (
	"projects/entitys"
	"projects/repositorys/authorRepository"
)

func AddUser(name string, email string, password string, image string) (error, entitys.User) {
	return authorRepository.AddUser(name, email, password, image)

}

func FindUserByEmail(email string) (error, entitys.User) {
	return authorRepository.FindUserByEmail(email)
}

func FindUserById(id int) (error, entitys.User) {
	return authorRepository.FindUserById(id)

}

func FindUserByUserName(name string) (error, entitys.User) {
	return authorRepository.FindUserByUserName(name)
}

func FindUserByUserEmail(name string) (error, entitys.User) {
	return authorRepository.FindUserByUserEmail(name)
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
