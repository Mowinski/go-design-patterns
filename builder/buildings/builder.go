package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type BuilderBuilding interface {
	BuildRooms() ([]Room, error)
	Build()	(Building, error)

	GetOwner() player.Player
}
