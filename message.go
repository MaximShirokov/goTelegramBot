// Copyright © 2019 Max Shirokov
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package gotelegrambot

type Message struct {
	client *Client

	// Key metadata
	MessageID int `json:"message_id"`
	Text 	  string `json:"text"`

	// Chat
	Chat *Chat
}

func (c *Client) SendMessage(chat *Chat, message *Message, extraParams RequestParams) error {
	path := "sendMessage"

	params := RequestParams{
		"chat_id": chat.ID,
		"text":    message.Text,
	}
	
	err := c.Post(path, params, &message)
	
	if err == nil {
		message.client = c
	}
	
	return err
}
