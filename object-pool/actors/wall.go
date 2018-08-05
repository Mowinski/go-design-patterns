package actors

import (
	"../interfaces"
)

type Wall struct {
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
