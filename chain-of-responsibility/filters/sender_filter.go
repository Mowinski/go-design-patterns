package filters

import (
	"strings"

	"../message"
)

type SenderFilter struct {
	next Filter
}

func (sf SenderFilter) Handle(messages []message.Message) {
	for index := range messages {
		isSpam := strings.Contains(strings.ToLower(messages[index].From), "spam")
		if isSpam {
			messages[index].Spam = true
		}
	}
	if sf.next != nil {
		sf.next.Handle(messages)
	}
}

func (sf *SenderFilter) AddNext(filter Filter) {
	sf.next = filter
}
