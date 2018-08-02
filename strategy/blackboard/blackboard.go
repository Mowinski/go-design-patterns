package blackboard

import (
	"../mob"
	"../player"
)

type BlackBoard struct {
	Mobs []mob.Mob
	Player player.Player

	isMobsTurn bool
}

func (bb* BlackBoard) AddMob(mob mob.Mob) {
	bb.Mobs = append(bb.Mobs, mob)
}

func (bb* BlackBoard) SetPlayer(player player.Player) {
	bb.Player = player
}

func (bb BlackBoard) PlayTurn() {
	// TODO Refactor it
	if bb.isMobsTurn {
		for _, mob := range bb.Mobs {
			mob.TakeAction()
		}
		return
	}
	bb.Player.TakeAction()
}