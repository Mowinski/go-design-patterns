package inbox

import (
	"../filters"
	"../message"
)

type Inbox struct {
	Messages []message.Message
	filters  []*filters.Filter
}

func NewInbox(fileName string) (Inbox, error) {
	var inbox Inbox
	messages, err := message.ReadMessagesFromJson(fileName)
	inbox.Messages = messages
	return inbox, err
}

func (i *Inbox) AddFilter(filter *filters.Filter) {
	if len(i.filters) != 0 {
		i.filters[len(i.filters)-1].AddNextFilter(filter)
	}
	i.filters = append(i.filters, filter)
}

func (i Inbox) Filter() {
	if len(i.filters) == 0 {
		return
	}
	i.filters[0].ValidateAllMessages(i.Messages)
}

func (i Inbox) DisplayAllMessagesOnConsole() {
	for _, msg := range i.Messages {
		msg.DisplayOnConsole()
	}
}
