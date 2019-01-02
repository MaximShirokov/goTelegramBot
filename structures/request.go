package structures

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func Post(url string, data interface{}) []byte {
	jsonStr, err := json.Marshal(data)
	
	if err != nil {
		panic(err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func Get(url string, data map[string]string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		panic(err.Error())
	}

	q := req.URL.Query()
	
	for key, value := range data {
		q.Add(key, value)
	}
	
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	return body
}