package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type HeatingBuilding struct {
	owner player.Player
}


func (hb HeatingBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = hb.BuildRoom()
	building.Owner = hb.GetOwner()
	building.Endurance = 350.99
	
	return building, err
}


func (hb HeatingBuilding) BuildRoom() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,		
	}
	
	ovenRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(10, 10), GetPoint(10, 15), GetPoint(0, 15)},
		WallColor: BLUE,
	}

	return []Room{mainRoom, ovenRoom}, nil
}

func (hb HeatingBuilding) GetOwner() player.Player {
	return hb.owner
}
