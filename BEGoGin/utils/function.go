package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

// Save uploaded file into upload folder
func SaveUploadedFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// Open the uploaded file
	srcFile, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	//Workaroud fix dupplicated name
	timeNow := time.Now().UnixNano() / int64(time.Millisecond)

	// Create the destination file on the server
	destinationPath := fmt.Sprintf("%s/%s", uploadDir, fmt.Sprintf("%d", timeNow)+"_"+fileHeader.Filename)
	dstFile, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// Copy the contents of the uploaded file to the destination file
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", err
	}
	//Temporarily get sv address from config
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	return "http://" + config.SVAddress + "/" + destinationPath, nil
}

// Marshal into JSON and return text
// eg: [1,2,3] => "[1,2,3]"

func MarshalGetString(data any) string {
	marshalJson, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(marshalJson)
}
