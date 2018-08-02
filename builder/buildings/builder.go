package buildings

import (
	"../player"
)

type BuilderBuilding interface {
	BuildRooms() ([]Room, error)
	Build()	(Building, error)

	GetOwner() player.Player
}
