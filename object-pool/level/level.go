package level

import (
	"fmt"
			"../interfaces"
		"../object_pool"
)

type Level struct {
	Board [][]rune

	XSize, YSize uint

	wallPool object_pool.WallPool
	actors []interfaces.Actor
	generatingWall bool
}

func NewLevel(XSize, YSize uint, wallPool object_pool.WallPool) Level {
	board := make([][]rune, XSize)
	for i := range board {
		board[i] = make([]rune, YSize)
	}

	level := Level{Board: board, XSize: XSize, YSize: YSize, wallPool: wallPool, generatingWall: true}
	level.clear()

	return level
}

func (l *Level) RenderOnConsole() {
	l.clear()
	l.addWall()
	l.renderActors()
	l.printRenderedBoardOnConsole()
}

func (l *Level) AddActorToRenderList(actor interfaces.Actor) {
	l.actors = append(l.actors, actor)
}

func (l *Level) SetObjectOnBoard(x, y uint, object rune) {
	l.Board[y][x] = object
}


func (l *Level) renderActors() {
	for _, actor := range l.actors {
		actor.Render(l)
	}
}

func (l *Level) printRenderedBoardOnConsole() {
	for x := range l.Board {
		for y := range l.Board[x] {
			fmt.Printf("%c", l.Board[x][y])
		}
		fmt.Printf("\n")
	}
}

func (l *Level) clear() {
	for x := range l.Board {
		for y := range l.Board[x] {
			l.Board[x][y] = ' '
		}
	}
}

func (l *Level) addWall() {
	
}