package storemedia

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"safe-cash-service/service/storemedia"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Image = "image"
	Video = "video"
)

//handleFormFile ...
func handleFormFile(c *gin.Context, fileNameRequest, userID, storeID string) error {
	file, header, err := c.Request.FormFile(fileNameRequest)
	if err != nil {
		if file == nil {
			return nil
		}
		return err
	}

	fileName := uuid.New().String()
	sourceNames := strings.Split(header.Filename, ".")
	tail := sourceNames[len(sourceNames)-1]

	// Save file here
	fileName = fileName + "." + tail
	out, err := os.Create("static/" + fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	typeMedia := ""
	thumbnail := ""
	if strings.HasPrefix(header.Header.Get("Content-type"), Image) {
		typeMedia = Image
		thumbnail = fileName
	}
	if strings.HasPrefix(header.Header.Get("Content-type"), Video) {
		typeMedia = Video
		thumbnail, err = generateThumbnailFromVideo(fileName)
		if err != nil {
			return err
		}
	}

	_, err = storemedia.CreateStoreMedia(fileName, typeMedia, thumbnail, &userID, &storeID)
	if err != nil {
		return err
	}

	return nil
}

func generateThumbnailFromImage(fileName string) {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	// Draw a red dot at (2, 3)
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})

	// Save to out.png
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func generateThumbnailFromVideo(fileName string) (string, error) {
	outputName := uuid.New().String()
	cmd := exec.Command("ffmpeg", "-i", "static/"+fileName, "-ss", "00:00:01.000", "-vframes", "1", "static/"+outputName+".png")
	if err := cmd.Run(); err != nil {
		return "", err
	}
	log.Println("fileName:", fileName)
	return fileName, nil
}

//Upload :
func Upload(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	err := handleFormFile(c, "snap_shot_desktop", userID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	err = handleFormFile(c, "snap_shot_webcam", userID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	err = handleFormFile(c, "video", userID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

//GetByStoreID :
func GetByStoreID(c *gin.Context) {

}
