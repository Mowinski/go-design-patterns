package main

import (
	"./mob"
	"./strategies"
)

func main() {
	supportMob := mob.Mob{Strategy: strategies.SupportStrategy{}}
	aggressiveMob := mob.Mob{Strategy: strategies.AggressiveStrategy{}}
	defensiveMob := mob.Mob{Strategy: strategies.DefensiveStrategy{}}


	supportMob.TakeAction()
	aggressiveMob.TakeAction()
	defensiveMob.TakeAction()
}
