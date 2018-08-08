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

	currentLevel interfaces.ILevel
}


func NewPlayer(x, y uint, level interfaces.ILevel) Player {
	return Player{X: x, Y: y, currentLevel: level}
}

func (p Player) Render(level interfaces.ILevel) {
	level.SetObjectOnBoard(p.X, p.Y, 'p')
}

func (p *Player) Move(direction Direction) {
	switch direction {
	case DOWN:
		p.tryMoveDown()

	case LEFT:
		p.tryMoveLeft()

	case RIGHT:
		p.tryMoveRight()
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


func (p *Player) tryMoveDown() {
	_, levelYSize := p.currentLevel.GetLevelSize()
	if p.Y >= levelYSize {
		return
	}
	p.Y += 1
}


func (p *Player) tryMoveLeft() {
	if p.X <= 0 {
		return
	}
	p.X -= 1
}


func (p *Player) tryMoveRight() {
	levelXSize, _ := p.currentLevel.GetLevelSize()
	if p.X >= levelXSize {
		return
	}
	p.X += 1
}