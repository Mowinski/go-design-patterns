package filters

import (
	"strings"

	"../message"
)

type TitleFilter struct {
	next Filter
}

func (tf TitleFilter) Handle(messages []message.Message) {
	for index := range messages {
		isSpam := strings.Contains(strings.ToLower(messages[index].Title), "spam")
		if isSpam {
			messages[index].Spam = true
		}
	}
	if tf.next != nil {
		tf.next.Handle(messages)
	}
}

func (tf *TitleFilter) AddNext(filter Filter) {
	tf.next = filter
}
