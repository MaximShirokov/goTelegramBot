package gotelegrambot

import (
	"github.com/MaximShirokov/goTelegramBot/structures"
)

func (c *config) SendMessage(msr MessageSendRequest) (int, error) {
	
	if msr.ChatID == "" {
		return 0, errors.New("Chat ID is empty")
	}

	if msr.Text == "" {
		return 0, errors.New("Text is empty")
	}

	url := c.URL + c.BotID + ":" + c.Token + "/sendMessage"
	resp, err := c.Post(url, msr)
	
	if err != nil {
		return 0, err
	}
	
	return c.GetResponseID(resp)
}
