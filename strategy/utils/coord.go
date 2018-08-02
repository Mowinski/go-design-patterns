package utils


type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST

)
type Coord struct {
	X int
	Y int
}

