package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

//getFileTypeAndTail ...
func getFileTypeAndTail(fileHandler *multipart.FileHeader) (string, string, error) {

	contentType := fileHandler.Header.Get("Content-Type")
	//typeFile/tailFile

	temp := strings.Split(contentType, "/")

	if len(temp) != 2 {
		return "", "", errors.New("Content type is not valid")
	}

	typeFile := temp[0]
	tailFile := temp[1]

	return typeFile, tailFile, nil
}

//Upload ...
func Upload(w http.ResponseWriter, r *http.Request) error {

	//Separate function to 3 parts:
	// - Read file upload from client
	// - Generate file name from client
	// - Save file upload to server

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `file`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return err
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileType, fileTail, _ := getFileTypeAndTail(handler)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	//

	log.Println(fileType)
	log.Println(fileTail)

	//tempFile, err := ioutil.TempFile("static/"+fileType, utils.GenerateUUID()+"."+fileTail)
	tempFile, err := ioutil.TempFile("static/"+fileType, "upload-*."+fileTail)

	log.Println(tempFile.Name())

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!

	//fmt.Fprintf(w, "Successfully Uploaded File\n")

	return nil
}
