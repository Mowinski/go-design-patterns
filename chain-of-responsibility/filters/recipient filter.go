package filters

import (
	"strings"

	"../message"
)

type RecipientFilter struct {
	next Filter
}

func (rf RecipientFilter) Handle(messages []message.Message) {
	for index := range messages {
		isSpam := strings.Compare(strings.ToLower(messages[index].To), "my_mail@example.com")
		if isSpam != 0 {
			messages[index].Spam = true
		}
	}
	if rf.next != nil {
		rf.next.Handle(messages)
	}
}

func (rf *RecipientFilter) AddNext(filter Filter) {
	rf.next = filter
}
