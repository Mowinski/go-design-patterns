package main

import (
	"fmt"
	"os"

	"./director"
	"./player"
)


func main() {
	playerOne := player.NewPlayer("TestPlayer")
	mapDirector, err := director.NewMapDirector(100, 100, playerOne)
	if err != nil {
		fmt.Printf("Error occure: %v", err)
		os.Exit(1)
	}
	mapCharacters := mapDirector.Level.Render()
	mapCharacters.ShowInConsole()
	err = mapCharacters.SaveToFile("map.jpeg", mapDirector.Level.SizeX, mapDirector.Level.SizeY)

	if err != nil {
		fmt.Printf("Error occure: %v", err)
		os.Exit(2)
	}
}
