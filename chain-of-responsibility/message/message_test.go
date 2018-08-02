package message_test

import (
	"../message"
	"testing"
)

func AssertStringEqual(t *testing.T, fieldName, result, expected string) {
	if result != expected {
		t.Errorf("Wrong %s field. Expected '%s', got '%s'", fieldName, expected, result)
	}
}

func TestReadMessagesFromJson(t *testing.T) {
	exampleFile := "examples_messages.json"

	messages, err := message.ReadMessagesFromJson(exampleFile)

	if err != nil {
		t.Fatalf("Error occurs: %v", err)
	}

	if len(messages) != 5 {
		t.Errorf("Wrong number of messages. Expected 5 got %d", len(messages))
	}

	firstMessage := messages[0]

	AssertStringEqual(t, "title", firstMessage.Title, "Test message")
	AssertStringEqual(t, "body", firstMessage.Body, "Lorem ipsum")
	AssertStringEqual(t, "date", firstMessage.Date, "2018-07-20 09:38:12")
	AssertStringEqual(t, "from", firstMessage.From, "example@example.com")
	AssertStringEqual(t, "to", firstMessage.To, "my_mail@example.com")
}
