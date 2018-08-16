package filters

import (
	"strings"

	"../message"
)

type TitleFilter struct {
	Validator
}

func (TitleFilter) IsSpam(msg message.Message) bool {
	return strings.Contains(strings.ToLower(msg.Title), "spam")
}
