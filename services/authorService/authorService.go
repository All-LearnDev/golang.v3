package authorService

import (
	"projects/entitys"
	"projects/repositorys/authorRepository"
	"projects/utils"
)

/**
 *  Insert user into Database.
 *
 * @param  email email of user
 * @param  password password of user
 * @param  name name  of user
 * @param  image image of user
 * @return  user or error when connection to database fail
 */
func InsertUser(name string, email string, password string, image string) (error, entitys.User) {
	var hash string
	hash, _ = utils.HashPassword(password)
	user := entitys.User{
		Name:     name,
		Email:    email,
		Password: hash,
		Image:    image,
	}
	return authorRepository.InsertUser(user)
}

/**
 *  Get user by email.
 *
 * @param  user email
 * @return  user or error when connection to database fail
 */
func GetUserByEmail(email string) (error, entitys.User) {
	return authorRepository.GetUserByEmail(email)
}

/**
 *  Get user by Id.
 *
 * @param  user id
 * @return  user or error when connection to database fail
 */
func GetUserById(id int) (error, entitys.User) {
	return authorRepository.GetUserById(id)

}

/**
 *  Get user by name.
 *
 * @param  user name
 * @return  user or error when connection to database fail
 */
func GetUserByName(name string) (error, entitys.User) {
	return authorRepository.GetUserByName(name)
}

/**
 *  Get RefreshToken by user id.
 *
 * @param  user id
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByUserId(Id int) (error, entitys.RefreshToken) {
	return authorRepository.GetRefreshTokenByUserId(Id)

}

/**
 *  Get RefreshToken by user name.
 *
 * @param  user name
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByUserName(userName string) entitys.RefreshToken {

	return authorRepository.GetRefreshTokenByUserName(userName)

}

/**
 *  Get RefreshToken by jwt token.
 *
 * @param jwt token.
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByToken(token string) (error, entitys.RefreshToken) {
	return authorRepository.GetRefreshTokenByToken(token)

}

/**
 *  Insert RefreshToken into Database.
 *
 * @param  refreshToken refreshToken
 * @return  RefreshToken or error when connection to database fail
 */
func InsertRefreshToken(refreshToken entitys.RefreshToken) (error, entitys.RefreshToken) {
	return authorRepository.InsertRefreshToken(refreshToken)

}
