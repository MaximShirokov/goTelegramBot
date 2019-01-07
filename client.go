package gotelegrambot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	APIURL = "https://api.telegram.org"
)

type (
	config struct {
		URL string
		BotID string
		Token string
	}
)

func New(botID string, token string) (*config, error) {
	var err error

	if botID == "" {
		return nil, errors.New("botID is empty")
	}
	
	if token == "" {
		return nil, errors.New("token is empty")
	}
	
	conf := &config{
		URL: APIURL,
		BotID: botID,
		Token: token
	}
	
	_, err = url.Parse(APIURL)
	
	if err != nil {
		return nil, err
	}
	
	return conf, nil
}

func (c *config) Post(url string, data string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	return client.Do(req)
}
