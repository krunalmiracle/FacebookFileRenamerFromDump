package utils

import (
	"fmt"
	"strings"
	"time"

	MessagesModel "github.com/mrnegativetw/FacebookArchiveRenamer/models/messages"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Viewer struct{}

func (v Viewer) PrintMessageDetails(messages MessagesModel.Messages) {
	var photosUri string
	for i := 0; i < len(messages.Messages); i++ {
		senderName := Viewer{}.encodeToHumanReadable(messages.Messages[i].SenderName)
		timestamp := Viewer{}.convertTimestampMsToDateTime(messages.Messages[i].TimestampMs)
		content := Viewer{}.encodeToHumanReadable(messages.Messages[i].Content)

		fmt.Printf("%s <%s> %s", timestamp, senderName, content)

		if len(messages.Messages[i].Photos) != 0 {
			for j := 0; j < len(messages.Messages[i].Photos); j++ {
				photosUri = messages.Messages[i].Photos[j].Uri
				fileName := strings.Split(photosUri, "/")
				// fileName[4]
				fmt.Printf("%s%s/%s", "target/", fileName[3], fileName[4])
			}
		}

		fmt.Printf("\n")
	}
}

func (v Viewer) PrintMessage(messages MessagesModel.Messages) {
	for i := 0; i < len(messages.Messages); i++ {
		content := Viewer{}.encodeToHumanReadable(messages.Messages[i].Content)
		fmt.Printf("%s\n", content)
	}
}

func (v Viewer) encodeToHumanReadable(content string) string {
	toLatin := charmap.ISO8859_1.NewEncoder()
	inLatin, _, _ := transform.String(toLatin, content)
	return inLatin
}

func (v Viewer) convertTimestampMsToDateTime(timestampMs int) string {
	parsedTime := time.UnixMilli(int64(timestampMs))
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		parsedTime.Year(), parsedTime.Month(), parsedTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())
}
