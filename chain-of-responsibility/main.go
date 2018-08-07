package main

import (
	"fmt"
	"os"

	"./filters"
	"./inbox"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Necessary parameter <file_name> is missing")
		return
	}
	messagesFile := os.Args[1]

	incomingMessageBox, err := inbox.NewInbox(messagesFile)
	if err != nil {
		panic(err)
	}

	incomingMessageBox.AddFilter(&filters.TitleFilter{})
	incomingMessageBox.AddFilter(&filters.BodyFilter{})
	incomingMessageBox.AddFilter(&filters.RecipientFilter{})
	incomingMessageBox.AddFilter(&filters.SenderFilter{})

	incomingMessageBox.Filter()

	incomingMessageBox.DisplayAllMessagesOnConsole()
}
