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
	"safe-cash-service/display"
	"safe-cash-service/models"
	"safe-cash-service/service/storemedia"
	"safe-cash-service/service/unlockinglog"
	"safe-cash-service/service/user"
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
func handleFormFile(c *gin.Context, fileNameRequest, storeID, unlockingLogID string) (*models.StoreMedia, error) {
	file, header, err := c.Request.FormFile(fileNameRequest)
	if err != nil {
		if file == nil {
			return nil, nil
		}
		return nil, err
	}

	fileName := uuid.New().String()
	sourceNames := strings.Split(header.Filename, ".")
	tail := sourceNames[len(sourceNames)-1]

	// Save file here
	fileName = fileName + "." + tail
	out, err := os.Create(FileDir + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return nil, err
	}

	unlockingLog, err := unlockinglog.GetUnlockingLogByID(unlockingLogID)
	if err != nil {
		return nil, err
	}
	var userID *string
	if unlockingLog.ID != "" {
		userID = new(string)
		*userID = *unlockingLog.UserID
	}

	typeMedia := ""
	thumbnail := ""

	tail = strings.ToLower(tail)

	if strings.HasPrefix(header.Header.Get("Content-type"), Image) ||
		tail == "png" ||
		tail == "jpeg" ||
		tail == "jpg" {
		typeMedia = Image
		thumbnail = fileName
	}

	if strings.HasPrefix(header.Header.Get("Content-type"), Video) ||
		tail == "mp4" ||
		tail == "avi" {
		typeMedia = Video
		thumbnail, err = generateThumbnailFromVideo(fileName)
		if err != nil {
			return nil, err
		}

		media, err := storemedia.CreateStoreMedia(thumbnail, ThumbnailVideo, thumbnail, userID, &storeID)
		if err != nil {
			return nil, err
		}

		_, err = storemedia.CreateMediaUnlockingLog(&media.ID, &unlockingLogID)
		if err != nil {
			return nil, err
		}
	}

	media, err := storemedia.CreateStoreMedia(fileName, typeMedia, thumbnail, userID, &storeID)
	if err != nil {
		return nil, err
	}

	_, err = storemedia.CreateMediaUnlockingLog(&media.ID, &unlockingLogID)
	if err != nil {
		return nil, err
	}

	return &media, nil
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
	cmd := exec.Command("ffmpeg", "-i", FileDir+"/"+fileName, "-ss", "00:00:01.000", "-vframes", "1", FileDir+"/"+outputName+".png")
	if err := cmd.Run(); err != nil {
		return "", err
	}
	// log.Println("fileName:", fileName)
	return outputName + ".png", nil
}

//Upload :
func Upload(c *gin.Context) {

	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	unlockingLogID, ok := c.GetPostForm("unlocking_log_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	if _, err := os.Stat(FileDir); os.IsNotExist(err) {
		os.Mkdir(FileDir, 0700)
	}
	_, err := handleFormFile(c, "screenshot", storeID, unlockingLogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	_, err = handleFormFile(c, "video", storeID, unlockingLogID)
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
	interfaceStoreID, _ := c.Get("store_id")
	storeID := fmt.Sprintf("%v", interfaceStoreID)

	date := c.Query("date")

	unLockingLogs, err := unlockinglog.GetUnlockingLogsByStoreID(storeID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	storeMediasDisplay := []display.StoreMedia{}

	for _, v := range unLockingLogs {
		storeMediaDisplay := display.StoreMedia{}
		storeMediaDisplay.Images = []display.Image{}
		storeMediaDisplay.Videos = []display.Video{}

		mediaUnlockingLog, err := storemedia.GetMediaUnlockingLogByUnlockingLogID(v.ID)
		if err != nil {
			continue
		}
		for _, value := range mediaUnlockingLog {
			storeMedia, err := storemedia.GetByID(*value.StoreMediaID)
			if err != nil {
				continue
			}

			log.Println("storeMedia.Name:", storeMedia.Name)

			if storeMedia.Type == Video {
				video := display.Video{
					URL: Host + storeMedia.Name,
				}
				storeMediaDisplay.Videos = append(storeMediaDisplay.Videos, video)
			} else {
				if storeMedia.Type == Image || storeMedia.Type == ThumbnailVideo {
					img := display.Image{
						URL:  Host + storeMedia.Name,
						Type: storeMedia.Type,
					}
					storeMediaDisplay.Images = append(storeMediaDisplay.Images, img)
				}
			}

			// storeMediaDisplay.Medias = append(storeMediaDisplay.Medias, storeMedia)
		}
		storeMediaDisplay.UnlockID = v.ID
		storeMediaDisplay.Method = v.Method
		storeMediaDisplay.CreatedAt = v.CreatedAt
		storeMediaDisplay.UpdatedAt = v.UpdatedAt
		storeMediaDisplay.IsSuccess = *v.IsSuccess

		if v.UserID != nil {
			userInfo, err := user.GetUserByID(*v.UserID)
			if err != nil {
				continue
			}
			storeMediaDisplay.Username = userInfo.FirstName + " " + userInfo.LastName
		} else {
			storeMediaDisplay.Username = "Ai đó"
		}

		storeMediasDisplay = append(storeMediasDisplay, storeMediaDisplay)
	}

	c.JSON(http.StatusOK, storeMediasDisplay)
}

//GetByUnlockingID ...
func GetByUnlockingID(c *gin.Context) {
	unlockingLogID := c.Param("unlock_id")
	mediaUnlockingLog, err := storemedia.GetMediaUnlockingLogByUnlockingLogID(unlockingLogID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	storeMedias := []models.StoreMedia{}
	for _, v := range mediaUnlockingLog {
		storeMedia, err := storemedia.GetByID(*v.StoreMediaID)
		if err != nil {
			continue
		}
		storeMedia.Name = Host + storeMedia.Name
		storeMedias = append(storeMedias, storeMedia)
	}

	c.JSON(http.StatusOK, storeMedias)

}
