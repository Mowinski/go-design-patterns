package mob

import (
	"../strategies"
	"../utils"
)

type Mob struct {
	Strategy strategies.BehaviourStrategy

	Damage int
	Health int
	MovePerTurn int

	Position utils.Coord
	Direction utils.Direction
}


func (m Mob) TakeAction() {
	m.Strategy.TakeAction()
}