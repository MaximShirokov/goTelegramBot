// Copyright © 2019 Max Shirokov
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package gotelegrambot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/pkg/errors"
)

const DEFAULT_BASEURL = "https://api.telegram.org/"

type Client struct {
	Client   *http.Client
	Logger   logger
	BaseURL  string
	BotID    string
	Token    string
	testMode bool
	ctx      context.Context
}

type logger interface {
	Debugf(string, ...interface{})
}

func NewClient(botID, token string) *Client {
	return &Client{
		Client:   http.DefaultClient,
		BaseURL:  DEFAULT_BASEURL,
		BotID:    botID,
		Token:    token,
		testMode: false,
		ctx:      context.Background(),
	}
}

func (c *Client) Post(path string, requestParams RequestParams, target interface{}) error {
	params := requestParams.ToURLValues()
	
	c.log("[telegram] POST %s?%s", path, params.Encode())

	if c.BotID != "" {
		params.Set("key", c.BotID)
	}

	if c.Token != "" {
		params.Set("token", c.Token)
	}

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, nil)
	
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errors.Wrapf(err, "HTTP Read error on response for %s", url)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	err = decoder.Decode(target)
	
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s:\n%s", url, string(b))
	}

	return nil
}

func (c *Client) log(format string, args ...interface{}) {
	if c.Logger != nil {
		c.Logger.Debugf(format, args...)
	}
}
