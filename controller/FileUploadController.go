package controller

import (
	"net/http"
	"projects/entitys"
	"projects/utils"

	"github.com/labstack/echo/v4"
)

func UploadFiles(c echo.Context) error {
	var user entitys.User
	user.Name = c.FormValue("name")
	user.Email = c.FormValue("email")

	println(" Name ", user.Name, " email ", user.Email)
	var images []entitys.Images
	err, images := utils.Upload(c)
	if err != nil {
		println(" How to handle one to many : ", len(images))
	}
	//user.Images = images
	return c.JSON(http.StatusOK, "hi")
}
