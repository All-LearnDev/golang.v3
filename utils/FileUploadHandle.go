package utils

import (
	"io"
	"os"
	"path/filepath"
	"projects/entitys"

	"github.com/labstack/echo/v4"
)

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
		dst, err := os.Create(filepath.Join("D:/go/images", filepath.Base(file.Filename)))
		if err != nil {
			return err, nil
		}
		defer dst.Close()
		println(" dress " + dst.Name())
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err, nil
		}

	}
	return nil, images
}
