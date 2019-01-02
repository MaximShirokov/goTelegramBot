package structures

import (
	"encoding/json"

	"github.com/MaximShirokov/goTelegramBot/entities"
)

type Message struct {
	Config Configure
}

// {"ok":true,
// 	"result":
// 		{"message_id":8,
// 		"from":
// 		{"id":709666023,"is_bot":true,"first_name":"ComeinService","username":"ComeinServiceBot"},
// 		"chat":{"id":234568990,"first_name":"Maxim","last_name":"Shirokov","username":"MaxRaynor","type":"private"},
// 		"date":1546448360,"text":"Happy New Year!"}}

func (m *Message) Send(message entity.Message) entity.Message {
	res := DoPost(m.Config.URL + "/set", requestRoot)
	response := &entity.ContactsSetResponseRoot{}
	err := json.Unmarshal(res, response)
	
	if err != nil {
		panic(err.Error())
	}
	
	if contact.Id == 0 {
		for _, add := range response.Response.Contacts.Add {
			contact.Id = add.Id
		}
	}

	return contact
}

func (c *Contacts) add(contact entity.Contact) entity.ContactsSetRequestRoot {
	requestRoot := entity.ContactsSetRequestRoot{}
	requestRoot.Request.Contacts.Add = append(requestRoot.Request.Contacts.Add, contact)
	return requestRoot
}