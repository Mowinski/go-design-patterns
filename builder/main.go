package main

import (
	"fmt"
	"os"

	"github.com/Mowinski/go-design-patterns/builder/level"
	"github.com/Mowinski/go-design-patterns/builder/buildings"
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type builderErrorHandler struct {
	err error
}

func main() {
	var brHandler builderErrorHandler
	playerOne := player.NewPlayer("TestPlayer")

	heating := brHandler.build(buildings.HeatingBuilding{Owner: playerOne, X: 0, Y: 0})
	smith := brHandler.build(buildings.SmithBuilding{Owner: playerOne , X: 15, Y: 15})
	shop := brHandler.build(buildings.ShopBuilding{Owner: playerOne, X: 50, Y: 50})
	candyShop := brHandler.build(buildings.ShopBuilding{Owner: playerOne, X: 0, Y: 50})

	if brHandler.err != nil {
		fmt.Printf("Error occures durring building level %v", brHandler.err)
		os.Exit(1)
	}

	buildingsInMaps := []buildings.Building{heating, smith, shop, candyShop}
	firstMap := level.NewLevel(100, 100, buildingsInMaps)
	mapCharacters := firstMap.Render()
	mapCharacters.ShowInConsole()
	err := mapCharacters.SaveToFile("map.jpeg", firstMap.SizeX, firstMap.SizeY)
	if err != nil {
		fmt.Printf("Error occure: %v", err)
		os.Exit(2)
	}
}


func (e builderErrorHandler) build(builder buildings.BuilderBuilding) buildings.Building  {
	var building buildings.Building
	if e.err != nil {
		return building
	}

	building, e.err = builder.Build()
	return building
}
