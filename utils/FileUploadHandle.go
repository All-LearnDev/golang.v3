package utils

import (
	"io"
	"os"
	"path/filepath"
	"projects/entitys"

	"github.com/labstack/echo/v4"
)

const filePath = "D:/go/Store"

func Upload(c echo.Context) (error, []entitys.Images) {

	var images []entitys.Images
	form, err := c.MultipartForm()
	if err != nil {
		return err, nil
	}
	files := form.File["files"]
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err, nil
		}
		defer src.Close()
		var myfile entitys.Images
		myfile.Filename = file.Filename
		images = append(images, myfile)
		dst, err := os.Create(filepath.Join(filePath, filepath.Base(file.Filename)))
		if err != nil {
			return err, nil
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err, nil
		}

	}
	return nil, images
}

func SingleFileUpload(c echo.Context) (error, string) {
	var imageName string = ""
	// Source
	file, err := c.FormFile("image")
	if err != nil {
		return err, imageName
	}
	src, err := file.Open()
	if err != nil {
		return err, imageName
	}
	defer src.Close()
	// Destination
	imageName = file.Filename
	dst, err := os.Create(filepath.Join(filePath, filepath.Base(file.Filename)))
	if err != nil {
		return err, imageName
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err, imageName
	}
	return nil, imageName
}
