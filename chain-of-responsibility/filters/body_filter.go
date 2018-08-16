package filters

import (
	"strings"

	"../message"
)

type BodyFilter struct {
	Validator
}

func (BodyFilter) IsSpam(msg message.Message) bool {
	return strings.Contains(strings.ToLower(msg.Body), "spam")
}
