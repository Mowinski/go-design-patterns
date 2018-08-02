package message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
	From  string `json:"from"`
	To    string `json:"to"`
	Spam  bool
}

func ReadMessagesFromJson(fileName string) ([]Message, error) {
	var messages []Message
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		return messages, err
	}

	json.Unmarshal(byteValue, &messages)
	return messages, nil
}

func (m Message) DisplayOnConsole() {
	spamMark := ""
	if m.Spam {
		spamMark = "<SPAM>"
	}
	fmt.Printf(
		"%sTitle: %s\n   Date: %s\n   From <%s>\n   To: %s\n   Body: %s\n",
		spamMark, m.Title, m.Date, m.From, m.To, m.Body,
	)
}
