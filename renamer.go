package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	MessagesModel "github.com/mrnegativetw/FacebookArchiveRenamer/models/messages"
	Utils "github.com/mrnegativetw/FacebookArchiveRenamer/utils"
)

func getOriginalFileName(uri string) string {
	fileName := strings.Split(uri, "/")
	return fileName[4]
}

func getFileExtensionFromFileName(fileName string) string {
	return strings.Split(fileName, ".")[1]
}

func convertUnixTimestampToDateTime(fileCreationTimestamp int) string {
	parsedTime := time.Unix(int64(fileCreationTimestamp), 0)
	return fmt.Sprintf("%d%02d%02d_%02d%02d%02d",
		parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())
}

// [OK, but duplicated] File not foun.
func renameFiles(originalFileName string, creationTimestamp int, filesFolderPath string) {
	fmt.Printf("originalFileName: %s\n", originalFileName)
	fmt.Printf("with extension: %s\n", getFileExtensionFromFileName(originalFileName))

	originalPath := fmt.Sprintf("%s%s%s",
		baseFolderPath,
		filesFolderPath,
		originalFileName)

	newFileName := convertUnixTimestampToDateTime(creationTimestamp)
	newPath := fmt.Sprintf("%s%s%s.%s",
		baseFolderPath,
		filesFolderPath,
		newFileName,
		getFileExtensionFromFileName(originalFileName))

	// Check is file name duplicated, if so add timestamp by 1 sec.
	for Utils.IsFileExist(newPath) {
		creationTimestamp += 1
		newFileName = convertUnixTimestampToDateTime(creationTimestamp)
		newPath = fmt.Sprintf("%s%s%s.%s",
			baseFolderPath,
			filesFolderPath,
			newFileName,
			getFileExtensionFromFileName(originalFileName))
	}

	// Rename File.
	if Utils.IsFileExist(originalPath) {
		e := os.Rename(originalPath, newPath)
		fmt.Printf("[OK] %s\n", newFileName)
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Printf("[Not Found] %s\n", originalPath)
	}
}

func renamePhotosFromSingleJsonFile(messages MessagesModel.Messages) {
	// Loop through all messages.
	for i := 0; i < len(messages.Messages); i++ {

		// Check message type is photo
		if len(messages.Messages[i].Photos) != 0 {
			// Loop through photos, sometimes a message has more than one photo.
			for j := 0; j < len(messages.Messages[i].Photos); j++ {
				// Passing original photo name and creation timestamp to rename
				// photos.
				renameFiles(
					getOriginalFileName(messages.Messages[i].Photos[j].Uri),
					messages.Messages[i].Photos[j].CreationTimestamp, photosFolderPath)
			}
		}
		// Check message type is video
		if len(messages.Messages[i].Videos) != 0 {
			// Loop through photos, sometimes a message has more than one photo.
			for j := 0; j < len(messages.Messages[i].Videos); j++ {
				// Passing original photo name and creation timestamp to rename
				// photos.
				renameFiles(
					getOriginalFileName(messages.Messages[i].Videos[j].Uri),
					messages.Messages[i].Videos[j].CreationTimestamp, videosFolderPath)
			}
		}
		// Check message type is audio
		if len(messages.Messages[i].Audios) != 0 {
			// Loop through photos, sometimes a message has more than one photo.
			for j := 0; j < len(messages.Messages[i].Audios); j++ {
				// Passing original photo name and creation timestamp to rename
				// photos.
				renameFiles(
					getOriginalFileName(messages.Messages[i].Audios[j].Uri),
					messages.Messages[i].Audios[j].CreationTimestamp, audiosFolderPath)
			}
		}
	}
}

func renamePhotosFromAllJsonFile() {
	jsonFileCount := 1

	filePath := fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)

	// Loop through all json files.
	for Utils.IsFileExist(filePath) {
		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s has opend successfully!\n", filePath)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var messages MessagesModel.Messages
		json.Unmarshal(byteValue, &messages)

		renamePhotosFromSingleJsonFile(messages)

		jsonFileCount++
		filePath = fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)
	}
}
