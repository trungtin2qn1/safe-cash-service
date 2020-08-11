package user

import (
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Image          = "image"
	Video          = "video"
	Screenshot     = "screenshot"
	ThumbnailVideo = "thumbnail_video"
	FileDir        = "static"
	Host           = "http://35.240.249.239:5000/public/"
	// Host    = "http://localhost:5000/public"
)

//handleFormFile ...
func handleFormFile(c *gin.Context, fileNameRequest string) (string, error) {
	file, header, err := c.Request.FormFile(fileNameRequest)
	if err != nil {
		if file != nil {
			return "", nil
		}
		return "", err
	}

	fileName := uuid.New().String()
	sourceNames := strings.Split(header.Filename, ".")
	tail := sourceNames[len(sourceNames)-1]

	// Save file here
	fileName = fileName + "." + tail
	out, err := os.Create(FileDir + "/" + fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
