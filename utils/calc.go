package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	MessagesModel "github.com/mrnegativetw/FacebookArchiveRenamer/models/messages"
)

type Calculator struct{}

func (c Calculator) CalculateTotalMessage(baseFolderPath string) int {
	totalMessages := 0
	jsonFileCount := 1

	filePath := fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)

	for IsFileExist(filePath) {
		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		} else {
			// fmt.Printf("%s opened successfully!", filePath)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var messages MessagesModel.Messages
		json.Unmarshal(byteValue, &messages)

		// fmt.Printf("%d messages.\n", len(messages.Messages))

		totalMessages += len(messages.Messages)

		jsonFileCount++
		filePath = fmt.Sprintf("%smessage_%d.json", baseFolderPath, jsonFileCount)
	}

	return totalMessages
}
