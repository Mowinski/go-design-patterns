package main

type Message struct {
	Title string
	Body  string
	Date  string
	From  string
	To    string
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

func ReadMessagesFromJson(directory string) []Message {
	return []Message{}
}
