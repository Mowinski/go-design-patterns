package filters

import (
	"strings"

	"../message"
)

type RecipientFilter struct {
	Validator
}

func (RecipientFilter) IsSpam(msg message.Message) bool {
	return strings.Compare(strings.ToLower(msg.To), "my_mail@example.com") != 0
}
