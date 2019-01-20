// Copyright © 2019 Max Shirokov
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package gotelegrambot

type Chat struct {
	client     *Client
	ID         string `json:"id"`
	Title      string `json:"title"`
	Username   string `json:"username"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
}
