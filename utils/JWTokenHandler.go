package utils

import (
	"fmt"
	"net/http"
	"projects/entitys"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string          `json:"name"`
	Roles []entitys.Roles `json:"roles"`
	jwt.StandardClaims
}

func GenerateRefreshToken(userName string, roles []entitys.Roles) entitys.RefreshToken {
	// Set custom claims
	var refreshToken entitys.RefreshToken
	refreshToken.UserName = userName
	expiresAt := time.Now().Add(time.Hour * 12).Unix()
	refreshToken.ExpiresAt = expiresAt
	claims := &JwtCustomClaims{
		userName,
		roles,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="))
	refreshToken.Token = result
	if err != nil {
		return entitys.RefreshToken{}
	} else {
		return refreshToken
	}

}
func GenerateJWT(userName string, roles []entitys.Roles) string {
	// Set custom claims
	claims := &JwtCustomClaims{
		userName,
		roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	result, _ := token.SignedString([]byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="))

	return result

}

func ParseToken(c echo.Context) error {
	var userName string
	tokenString := c.Param("token")
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="), nil
	})
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		userName = claims.Name
	} else {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, userName)
}

func GetRolesFromToken(tokenString string) entitys.JUser {
	var user entitys.JUser
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="), nil
	})
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		user.Name = claims.Name
		user.Roles = claims.Roles
	} else {
		fmt.Println(err)
	}
	return user
}

func ValidToken(validToken string) bool {

	token, er := jwt.ParseWithClaims(validToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="), nil
	})
	if er != nil {
		return false
	}
	return token.Valid
}

func ExpiredToken(tokenString string) bool {
	var expireTime int64
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Y29uZHVvbmdzdWFlbWRpODkzNA=="), nil
	})
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		expireTime = claims.StandardClaims.ExpiresAt
	} else {
		fmt.Println(err)
	}
	now := time.Now().Unix()
	if now < expireTime {
		return true
	} else {
		return false
	}
}
