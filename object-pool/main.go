package main

import (
	"time"
	"math/rand"

	"./level"
	"./actors"
	"./object_pool"

	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	player := actors.NewPlayer(15, 0)
	wallPool := object_pool.NewWallPool()
	firstStage := level.NewLevel(30, 30, &wallPool)

	firstStage.AddActorToRenderList(&player)
	go wallPool.GenerateNewWall(firstStage.XSize, firstStage.YSize)
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