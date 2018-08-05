package actors

import (
	"fmt"

	"../interfaces"
)

type Direction rune

const (
	DOWN Direction = 's'
	LEFT Direction = 'a'
	RIGHT Direction = 'd'
)

type Player struct {
	X, Y uint
}


func NewPlayer(x, y uint) Player {
	return Player{X: x, Y: y}
}

func (p Player) Render(level interfaces.ILevel) {
	level.SetObjectOnBoard(p.X, p.Y, 'p')
}

func (p *Player) Move(direction Direction) {
	switch direction {
	case DOWN:
		p.Y += 1

	case LEFT:
		p.X -= 1

	case RIGHT:
		p.X += 1

	}
}

func (p *Player) TakeAction() Direction {
	var direction Direction
	for IsValidDirection(direction) != true {
		fmt.Print("> ")
		fmt.Scanf("%c", &direction)
	}
	return direction
}


func IsValidDirection(direction Direction) bool {
	isLeftMove := direction == LEFT
	isRightMove := direction == RIGHT
	isDownMove := direction == DOWN

	if !isLeftMove && !isRightMove && !isDownMove {
		return false
	}
	return true
}