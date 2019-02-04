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
	"net/url"
	"os"
	"github.com/pkg/errors"
)

const DEFAULT_BASEURL = "https://api.telegram.org"

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
	proxyURL, _ := url.Parse(os.Getenv("HTTP_PROXY"))
	httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	return &Client{
		Client:   httpClient,
		BaseURL:  DEFAULT_BASEURL,
		BotID:    botID,
		Token:    token,
		testMode: false,
		ctx:      context.Background(),
	}
}

func (c *Client) Get(path string, requestParams RequestParams, target interface{}) error {
	params := requestParams.ToURLValues()

	c.log("[telegram] GET %s?%s", path, params.Encode())

	if c.BotID == "" {
		return errors.New("BotID is empty")
	}

	if c.Token == "" {
		return errors.New("Token is empty")
	}

	url := fmt.Sprintf("%s/%s:%s/%s", c.BaseURL, c.BotID, c.Token, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}
	
	req = req.WithContext(c.ctx)

	resp, err := c.Client.Do(req)
	
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return makeHttpClientError(url, resp)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	
	if err != nil {
		return errors.Wrapf(err, "JSON decode failed on %s", url)
	}

	return nil
}

func (c *Client) Post(path string, requestParams RequestParams, target interface{}) error {
	params := requestParams.ToURLValues()
	
	c.log("[telegram] POST %s?%s", path, params.Encode())

	if c.BotID == "" {
		return errors.New("BotID is empty")
	}

	if c.Token == "" {
		return errors.New("Token is empty")
	}

	url := fmt.Sprintf("%s/%s:%s/%s", c.BaseURL, c.BotID, c.Token, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, nil)
	
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println(url)

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
