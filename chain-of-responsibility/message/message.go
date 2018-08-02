package message


import (
	"io/ioutil"
	"encoding/json"
)

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
	From  string `json:"from"`
	To    string `json:"to"`
}

func NewMessage(title, body, date, from, to string) Message {
	return Message{
		Title: title,
		Body:  body,
		Date:  date,
		From:  from,
		To:    to,
	}
}

func ReadMessagesFromJson(fileName string) ([]Message, error) {
	var messages []Message
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		return messages, nil
	}

	json.Unmarshal(byteValue, &messages)
	return messages, nil
}
