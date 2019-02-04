// Copyright © 2019 Max Shirokov
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package gotelegrambot

type Chat struct {
	client     *Client

	ID         string `json:"id"`
	Type       string `json:"type"`
 	Title      string `json:"title"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`

	// all_members_are_administrators Boolean
	// photo ChatPhoto

	Description string `json:"description"`
	InviteLink  string `json:"invite_link"`
	// pinned_message Message
	StickerSetName string `json:"sticker_set_name"`
	CanSetStickerSet bool `json:"can_set_sticker_set"`
}

func (c *Client) GetChat(chatID string, params RequestParams) (chat *Chat, err error) {
	path := "getChat"

	args := RequestParams{
		"chat_id": chatID,
	}
	
	err = c.Get(path, args, &chat)
	
	if chat != nil {
		chat.client = c
	}
	
	return
}