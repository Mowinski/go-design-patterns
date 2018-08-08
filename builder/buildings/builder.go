package buildings

import (
	"../player"
)

type IBuilding interface {
	Build()	(Building, error)

	getOwner() player.Player
	buildRooms() ([]Room, error)
}
