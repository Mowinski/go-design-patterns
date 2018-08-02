package actor

import (
	"../utils"
	"../mob"
)

type Actor interface {
	CanMove(target utils.Coord) bool
	CanShoot(target utils.Coord) bool

	Move(target utils.Coord)
	Shoot(target utils.Coord)
}

type MobActor interface {
	Actor

	Heal(mob mob.Mob)
}