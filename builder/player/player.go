package player

type Player struct {
	ID int
	Name string
}

func NewPlayer(name string) Player {
	var player Player
	player.ID = 123 // TODO randomize it
	player.Name = name
	return player
}
