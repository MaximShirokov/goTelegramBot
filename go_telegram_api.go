package gotelegrambotgo

import (
	"github.com/MaximShirokov/goTelegramBot/structures"
)

const (
	HTTPOK = 200
	APIURL = "https://api.telegram.org"
)

type TelegramAPI struct {
	Config structures.Configure
}

func New(botID, token string) TelegramApi {
	var err error

	config := structures.Configure{
		Url:   APIURL,
		BotID: botID,
		Token: token,
	}

	if err != nil {
		panic(err.Error())
	}

	telegramAPI := TelegramAPI{
		config,
	}

	return telegramAPI
}
