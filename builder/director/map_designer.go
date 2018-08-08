package director

import (
	"../buildings"
	"../level"
	"../player"
)

type MapDirector struct {
	Level level.Level

	err error
}


func NewMapDirector(sizeX, sizeY int, player player.Player) (MapDirector, error) {
	var mapDesigner MapDirector

	heating := mapDesigner.build(buildings.HeatingBuilding{Owner: player, X: 0, Y: 0})
	smith := mapDesigner.build(buildings.SmithBuilding{Owner: player , X: 15, Y: 15})
	shop := mapDesigner.build(buildings.ShopBuilding{Owner: player, X: 50, Y: 50})
	candyShop := mapDesigner.build(buildings.ShopBuilding{Owner: player, X: 0, Y: 50})

	buildingsInMaps := []buildings.Building{*heating, *smith, *shop, *candyShop}
	mapDesigner.Level = level.NewLevel(sizeX, sizeY, buildingsInMaps)

	return mapDesigner, mapDesigner.err
}


func (md MapDirector) build(builder buildings.IBuilding) *buildings.Building  {
	var building buildings.Building
	if md.err != nil {
		return nil
	}

	building, md.err = builder.Build()
	return &building
}
