package entity

type Message struct {
	Chat Chat `json:"chat"`
	Text string `json:"text"`
}
