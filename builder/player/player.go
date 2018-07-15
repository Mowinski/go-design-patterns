package player

import "math/rand"

type Player struct {
	ID int
	Name string
}

func NewPlayer(name string) Player {
	var player Player
	player.ID = rand.Int()
	player.Name = name
	return player
}
