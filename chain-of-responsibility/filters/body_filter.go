package filters

import (
	"../message"
	"strings"
)

type BodyFilter struct {
	next Filter
}

func (bf BodyFilter) Handle(messages []message.Message) {
	for index := range messages {
		isSpam := strings.Contains(strings.ToLower(messages[index].Body), "spam")
		if isSpam {
			messages[index].Spam = true
		}
	}
	if bf.next != nil {
		bf.next.Handle(messages)
	}
}

func (bf *BodyFilter) AddNext(filter Filter) {
	bf.next = filter
}
