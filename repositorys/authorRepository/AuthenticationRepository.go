package authorRepository

import (
	"projects/configs"
	"projects/entitys"
)

var Connection = configs.GetConnection()

/**
 *  Insert user into Database.
 *
 * @param  user info
 * @return  user or error when connection to database fail
 */
func InsertUser(user entitys.User) (error, entitys.User) {

	error := Connection.Create(&user).Error
	return error, user

}

/**
 *  Get user by email.
 *
 * @param  user email
 * @return  user or error when connection to database fail
 */
func GetUserByEmail(email string) (error, entitys.User) {
	var user entitys.User
	error := Connection.Where("Email = ?", email).First(&user).Error
	return error, user
}

/**
 *  Get user by Id.
 *
 * @param  user id
 * @return  user or error when connection to database fail
 */
func GetUserById(id int) (error, entitys.User) {
	var user entitys.User
	error := Connection.Where("Id = ?", id).First(&user).Error
	return error, user
}

/**
 *  Get user by name.
 *
 * @param  user name
 * @return  user or error when connection to database fail
 */
func GetUserByName(name string) (error, entitys.User) {
	var user entitys.User
	error := Connection.Where("Name = ?", name).First(&user).Error
	return error, user
}

/**
 *  Get RefreshToken by user id.
 *
 * @param  user id
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByUserId(Id int) (error, entitys.RefreshToken) {
	var token entitys.RefreshToken
	error := Connection.Where("Id = ?", Id).First(&token).Error
	return error, token
}

/**
 *  Get RefreshToken by user name.
 *
 * @param  user name
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByUserName(userName string) entitys.RefreshToken {
	var token entitys.RefreshToken
	Connection.First(&token, "UserName = ?", userName)
	return token
}

/**
 *  Get RefreshToken by jwt token.
 *
 * @param jwt token.
 * @return  RefreshToken or error when connection to database fail
 */
func GetRefreshTokenByToken(token string) (error, entitys.RefreshToken) {
	var refreshToken entitys.RefreshToken
	error := Connection.Where("Token = ?", token).First(&refreshToken).Error
	return error, refreshToken
}

/**
 *  Insert RefreshToken into Database.
 *
 * @param  refreshToken refreshToken
 * @return  RefreshToken or error when connection to database fail
 */
func InsertRefreshToken(refreshToken entitys.RefreshToken) (error, entitys.RefreshToken) {
	error := Connection.Create(&refreshToken).Error
	return error, refreshToken

}
