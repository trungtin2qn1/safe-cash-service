package storemedia

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"safe-cash-service/service/storemedia"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	Image = "image"
	Video = "video"
)

//Upload :
func Upload(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileName := uuid.New().String()

	sourceNames := strings.Split(header.Filename, ".")
	tail := sourceNames[len(sourceNames)-1]

	typeMedia := ""
	if strings.HasPrefix(header.Header.Get("Content-type"), Image) {
		typeMedia = Image
		// TODO: @tin process image here
	}
	if strings.HasPrefix(header.Header.Get("Content-type"), Video) {
		typeMedia = Video
		// TODO: @tin process video here
	}

	// Save file here
	fileName = fileName + "." + tail
	out, err := os.Create("static/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: @tin process thumbnail here

	media, err := storemedia.CreateStoreMedia(fileName, typeMedia, fileName, &userID, &storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't create store media",
		})
		return
	}

	c.JSON(200, media)
}

//GetByStoreID :
func GetByStoreID(c *gin.Context) {

}
