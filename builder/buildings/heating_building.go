package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type HeatingBuilding struct {
	Owner player.Player
	X int
	Y int
}


func (hb HeatingBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = hb.BuildRooms()
	building.Owner = hb.GetOwner()
	building.Endurance = 350.99
	building.X = hb.X
	building.Y = hb.Y
	
	return building, err
}


func (hb HeatingBuilding) BuildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,
		RoomLetter: 'm',
	}
	
	ovenRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(10, 10), GetPoint(10, 15), GetPoint(0, 15)},
		WallColor: BLUE,
		RoomLetter: 'o',
	}

	return []Room{mainRoom, ovenRoom}, nil
}

func (hb HeatingBuilding) GetOwner() player.Player {
	return hb.Owner
}
