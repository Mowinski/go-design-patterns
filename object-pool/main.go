package main

import (
	"math/rand"
	"time"

	"./actors"
	"./level"
	"./object_pool"

	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	wallPool := object_pool.NewWallPool()
	firstStage := level.NewLevel(30, 30, &wallPool)

	player := actors.NewPlayer(15, 0)
	firstStage.AddActorToRenderList(&player)

	go wallPool.GenerateNewWall(firstStage.XSize, firstStage.YSize)
	defer wallPool.StopGeneratingWalls()

	loadInitialCountOfWall(5)

	for i := 0; i < 20; i++ {
		direction := player.TakeAction()
		player.Move(direction)
		firstStage.RenderOnConsole()
	}

}

func loadInitialCountOfWall(count int) {
	rand.Seed(time.Now().Unix())
	numberOfTicks := count * 20
	bar := pb.StartNew(numberOfTicks)
	bar.ShowCounters = false

	defer bar.FinishPrint("Game loaded!")
	for i := 0; i < numberOfTicks; i++ {
		time.Sleep(time.Millisecond * 100)
		bar.Increment()
	}

}
