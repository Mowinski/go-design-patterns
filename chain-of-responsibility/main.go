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

	incomingMessageBox.AddFilter(filters.NewFilter(filters.TitleFilter{}))
	incomingMessageBox.AddFilter(filters.NewFilter(filters.BodyFilter{}))
	incomingMessageBox.AddFilter(filters.NewFilter(filters.RecipientFilter{}))
	incomingMessageBox.AddFilter(filters.NewFilter(filters.SenderFilter{}))

	incomingMessageBox.Filter()

	incomingMessageBox.DisplayAllMessagesOnConsole()
}
