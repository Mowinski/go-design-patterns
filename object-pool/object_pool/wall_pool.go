package object_pool

import (
	"../actors"
	"math/rand"
	"time"
)

const MAXSIZE = 5


type WallPool struct {
	walls map[uint][]actors.Wall

	generatingWall bool
}


func NewWallPool() WallPool {
	wallsMap := make(map[uint][]actors.Wall)
	return WallPool{walls:wallsMap, generatingWall: true}
}


func (wp *WallPool) GenerateNewWall(xLevelSize, yLevelSize uint) {
	for wp.generatingWall {
		randomX := uint(rand.Int31n(int32(xLevelSize - MAXSIZE)))
		randomY := uint(rand.Int31n(int32(yLevelSize)))
		randomWidth := uint(rand.Int31n(MAXSIZE))

		wall := actors.NewWall(randomX, randomY, randomWidth)
		wp.walls[randomWidth] = append(wp.walls[randomWidth], wall)

		time.Sleep(time.Second * 2)
	}
}