package filters

import (
	"strings"

	"../message"
)

type SenderFilter struct {
	Validator
}

func (SenderFilter) IsSpam(msg message.Message) bool {
	return strings.Contains(strings.ToLower(msg.From), "spam")
}
