package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("CHANNEL_ID", "C5X3GDU5D")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := os.Getenv("CHANNEL_ID")

	fileArr := []string{"HelloWorld.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileV2Parameters{
			Channel:  channel,
			Filename: fileArr[i],
		}
		file, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, Id: %s\n", file.Title, file.ID)
	}
}
