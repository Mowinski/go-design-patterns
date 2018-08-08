package buildings

import (
	"../player"
)

type HeatingBuilding struct {
	Owner player.Player
	X int
	Y int
}


func (hb HeatingBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = hb.buildRooms()
	building.Owner = hb.getOwner()
	building.Endurance = 350.99
	building.X = hb.X
	building.Y = hb.Y
	
	return building, err
}


func (hb HeatingBuilding) buildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		RoomLetter: 'm',
	}
	
	ovenRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(10, 10), GetPoint(10, 15), GetPoint(0, 15)},
		RoomLetter: 'o',
	}

	return []Room{mainRoom, ovenRoom}, nil
}

func (hb HeatingBuilding) getOwner() player.Player {
	return hb.Owner
}
