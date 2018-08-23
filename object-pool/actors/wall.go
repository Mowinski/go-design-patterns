package actors

import (
	"math/rand"
	"time"

	"../interfaces"
)

const MAXSIZE = 5


type Wall struct {
	interfaces.Actor

	X, Y uint
	Width uint
}


func (w Wall) Render(level interfaces.ILevel) {
	for i := uint(0); i < w.Width; i++ {
		level.SetObjectOnBoard(w.X + i, w.Y, '-')
	}
}


func NewWall(x, y, width uint) Wall {
	return Wall{X: x, Y: y, Width: width}
}


func GenerateRandomWidthWall(xLevelSize, yLevelSize uint) Wall {
	time.Sleep(time.Second * 2)

	randomX := uint(rand.Int31n(int32(xLevelSize - MAXSIZE)))
	randomY := uint(rand.Int31n(int32(yLevelSize)))
	randomWidth := uint(rand.Int31n(MAXSIZE + 1))

	return NewWall(randomX, randomY, randomWidth)
}
