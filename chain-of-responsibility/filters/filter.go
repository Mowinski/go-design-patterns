package filters

import (
	"sync"

	"../message"
)

type Validator interface {
	IsSpam(message.Message) bool
}

type Filter struct {
	validator Validator

	next *Filter
}

func NewFilter(filter Validator) *Filter {
	return &Filter{validator: filter, next: nil}
}

func (f Filter) ValidateAllMessages(messages []message.Message) {
	var wg sync.WaitGroup

	for index := range messages {
		wg.Add(1)
		go f.validateMessage(&messages[index], &wg)
	}
	wg.Wait()

	if f.isNextHandler() {
		f.next.ValidateAllMessages(messages)
	}
}

func (f *Filter) AddNextFilter(filter *Filter) {
	f.next = filter
}

func (f Filter) isNextHandler() bool {
	return f.next != nil
}

func (f Filter) validateMessage(msg *message.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	if f.validator.IsSpam(*msg) {
		msg.Spam = true
	}
}
