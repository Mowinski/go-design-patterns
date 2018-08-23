package object_pool

import (
	"../actors"
)


type WallPool struct {
	walls map[uint][]actors.Wall

	generatingWall bool
}

func NewWallPool() WallPool {
	wallsMap := make(map[uint][]actors.Wall)
	return WallPool{walls: wallsMap, generatingWall: true}
}

func (wp *WallPool) GenerateWalls(xLevelSize, yLevelSize uint) {
	for wp.generatingWall {
		wall := actors.GenerateRandomWidthWall(xLevelSize, yLevelSize)
		wp.walls[wall.Width] = append(wp.walls[wall.Width], wall)
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
