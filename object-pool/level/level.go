package level

import (
	"fmt"

	"../interfaces"
	"../object_pool"
	"../actors"
)

type Level struct {
	Board [][]rune

	XSize, YSize uint

	wallPool      *object_pool.WallPool
	actors        []interfaces.Actor
	nextWallWidth uint
}

func NewLevel(XSize, YSize uint, wallPool *object_pool.WallPool) Level {
	board := make([][]rune, XSize)
	for i := range board {
		board[i] = make([]rune, YSize)
	}

	level := Level{Board: board, XSize: XSize, YSize: YSize, wallPool: wallPool}
	level.clear()

	return level
}

func (l *Level) RenderOnConsole() {
	l.clear()
	l.addWall()
	l.renderActors()
	l.printWallStatsSummary()
	l.printRenderedBoardOnConsole()
}

func (l *Level) AddActorToRenderList(actor interfaces.Actor) {
	l.actors = append(l.actors, actor)
}

func (l *Level) SetObjectOnBoard(x, y uint, object rune) {
	l.Board[y][x] = object
}

func (l Level) GetLevelSize() (uint, uint) {
	return l.XSize, l.YSize
}

func (l *Level) renderActors() {
	for _, actor := range l.actors {
		actor.Render(l)
	}
}

func (l Level) printWallStatsSummary() {
	summary := "Object pool stats:\n Width 1 -> %d; Width 2 -> %d; Width 3 -> %d; Width 4 -> %d; Width 5 -> %d\n\n"
	fmt.Printf(
		summary,
		l.wallPool.GetWallByWidthCount(1),
		l.wallPool.GetWallByWidthCount(2),
		l.wallPool.GetWallByWidthCount(3),
		l.wallPool.GetWallByWidthCount(4),
		l.wallPool.GetWallByWidthCount(5),
	)
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
	l.nextWallWidth = (l.nextWallWidth + 1) % (actors.MAXSIZE + 1)
	wall := l.wallPool.GetInstanceWithWidth(l.nextWallWidth)

	l.AddActorToRenderList(&wall)
}
