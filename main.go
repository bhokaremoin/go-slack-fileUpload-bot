package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_CHANNEL_ID", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"temp.txt"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error While Uploading File := %v", err)
			return
		}
		fmt.Printf("File Uploaded\nName: %v,\tURL: %v", file.Name, file.URLPrivate)
	}
}
