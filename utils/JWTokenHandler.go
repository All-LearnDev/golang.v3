package utils

import (
	"fmt"
	"net/http"
	"projects/entitys"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateRefreshToken(userName string, role bool) (error, entitys.RefreshToken) {
	// Set custom claims
	var refreshToken entitys.RefreshToken
	refreshToken.UserName = userName
	expiresAt := time.Now().Add(time.Hour * 72).Unix()
	refreshToken.ExpiresAt = expiresAt
	claims := &JwtCustomClaims{
		userName,
		true,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte("konmeo12397"))
	refreshToken.Token = result
	if err != nil {
		return err, entitys.RefreshToken{}
	} else {
		return nil, refreshToken
	}

}
func GenerateJWT(userName string, role bool) (error, string) {
	// Set custom claims
	claims := &JwtCustomClaims{
		userName,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	result, err := token.SignedString([]byte("konmeo12397"))
	if err != nil {
		return err, ""
	} else {
		return nil, result
	}

}

func ParseToken(c echo.Context) error {
	var userName string
	tokenString := c.Param("token")
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("konmeo12397"), nil
	})
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		userName = claims.Name
	} else {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, userName)
}

func ValidToken(validToken string) bool {

	token, _ := jwt.ParseWithClaims(validToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("konmeo12397"), nil
	})

	return token.Valid
}

func ExpiredToken(tokenString string) bool {
	var expireTime int64
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("konmeo12397"), nil
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
