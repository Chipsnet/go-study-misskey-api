package main

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"

	"gopkg.in/ini.v1"
)

func main() {
	var input string

	config, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err.Error())
    }

	apiKey := config.Section("main").Key("api_key").String()
	
	fmt.Println("投稿する内容を入力してください！")
	fmt.Print(">> ")
	fmt.Scan(&input)

    fmt.Println(input)

	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(apiKey),
		misskey.WithBaseURL("https", "misskey.io", ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
	}

	res, err := client.Notes().Create(notes.CreateRequest{
		Text: core.NewString(input),
		Visibility: models.VisibilityPublic,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	log.Println(res.CreatedNote.ID)
}