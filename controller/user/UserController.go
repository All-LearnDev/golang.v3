package controller

import (
	"net/http"
	"projects/entitys"
	userService "projects/services/user"
	"projects/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	var user entitys.User
	user.Name = c.FormValue("name")
	user.Email = c.FormValue("email")
	var images []entitys.Images
	err, images := utils.Upload(c)
	if err != nil {
		println(" How to handle one to many : ", len(images))
	}
	user.Images = images
	result := userService.AddUser(user)
	return c.JSON(http.StatusOK, result)
}

func ListUser(c echo.Context) error {
	result := userService.ListUser()
	return c.JSON(http.StatusOK, result)
}

func FindByUserId(c echo.Context) error {
	id := c.Param("id")
	intVar, err := strconv.Atoi(id)
	if err == nil {
		user := userService.FindByUserId(intVar)
		return c.JSON(http.StatusOK, user)
	} else {
		return c.String(http.StatusOK, "Record not found ")
	}
}
