package filters

import (
	"../message"
)

type Filter interface {
	Handle([]message.Message)

	AddNext(Filter)
}
