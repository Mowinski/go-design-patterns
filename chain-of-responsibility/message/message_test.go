package message_test

import (
	"testing"
	"github.com/Mowinski/go-design-patterns/chain-of-responsibility/message"
)

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

	if firstMessage.Title != "Test message" {
		t.Errorf("Wrong title field, Expected 'Test message', got '%s'", firstMessage.Title)
	}

	if firstMessage.Body != "Lorem ipsum" {
		t.Errorf("Wrong body field, Expected 'Lorem ipsum', got '%s'", firstMessage.Body)
	}
}
