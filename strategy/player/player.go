package player

import (
	"../utils"
)

type Player struct {
	Damage int
	Health int
	MovePerTurn int

	Position utils.Coord
	Direction utils.Direction
}

func (p Player) CanShoot(target utils.Coord) bool {
	return false
}

func (p Player) CanMove(target utils.Coord) bool {
	return false
}

func (p Player) TakeAction() {

}