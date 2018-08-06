package object_pool

import (
	"math/rand"
	"time"

	"../actors"
)

const MAXSIZE = 5

type WallPool struct {
	walls map[uint][]actors.Wall

	generatingWall bool
}

func NewWallPool() WallPool {
	wallsMap := make(map[uint][]actors.Wall)
	return WallPool{walls: wallsMap, generatingWall: true}
}

func (wp *WallPool) GenerateNewWall(xLevelSize, yLevelSize uint) {
	for wp.generatingWall {
		randomX := uint(rand.Int31n(int32(xLevelSize - MAXSIZE)))
		randomY := uint(rand.Int31n(int32(yLevelSize)))
		randomWidth := uint(rand.Int31n(MAXSIZE + 1))

		wall := actors.NewWall(randomX, randomY, randomWidth)
		wp.walls[randomWidth] = append(wp.walls[randomWidth], wall)

		time.Sleep(time.Second * 2)
	}
}

func (wp WallPool) GetInstanceWithWidth(width uint) actors.Wall {
	wall := wp.walls[width][0]
	wp.walls[width] = wp.walls[width][1:]
	return wall
}

func (wp WallPool) GetWallByWidthCount(width uint) int {
	return len(wp.walls[width])
}

func (wp *WallPool) StopGeneratingWalls() {
	wp.generatingWall = false
}
