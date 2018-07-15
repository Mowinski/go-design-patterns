package level

import (
	"fmt"

	"github.com/Mowinski/go-design-patterns/builder/buildings"
)

type Level struct {
	SizeX int
	SizeY int

	Buildings []buildings.Building
}

type Representation [][]rune

func NewLevel(sizeX, sizeY int, buildings []buildings.Building) Level {
	var level Level
	level.SizeX = sizeX
	level.SizeY = sizeY
	level.Buildings = buildings

	return level
}

func (l Level) Render() Representation {
	level :=  make([][]rune, l.SizeX)
	for x := 0; x < l.SizeX; x++ {
		level[x] = make([]rune, l.SizeY)
		for y := range level[x] {
			level[x][y] = l.getLetter(x, y)
		}
	}
	return level
}


func (r Representation) ShowInConsole() {
	for _, row := range r {
		for _, value := range row {
			fmt.Printf("%c", value)
		}
		fmt.Println()
	}
}

func (l Level) getLetter(x, y int) rune {
	letter := ' '
	point := buildings.GetPoint(x, y)
	for _, build := range l.Buildings {
		for _, room := range build.Rooms {
			if build.IsInRoom(point, room) {
				return room.RoomLetter
			}
		}
	}
	return letter
}
