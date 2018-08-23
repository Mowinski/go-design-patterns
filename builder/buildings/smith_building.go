package buildings

import (
	"../player"
)

type SmithBuilding struct {
	IBuilding

	Owner player.Player
	X int
	Y int
}

func (sb SmithBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.buildRooms()
	building.Endurance = 900.00
	building.Owner = sb.getOwner()
	building.X = sb.X
	building.Y = sb.Y
	
	return building, err
}

func (sb SmithBuilding) buildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		RoomLetter: 'm',
	}
	
	smithRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(20, 0), GetPoint(20, 5), GetPoint(10, 5)},
		RoomLetter: 's',
	}

	anvilRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(5, 10), GetPoint(5, 15), GetPoint(0, 15)},
		RoomLetter: 'a',
	}
	
	return []Room{mainRoom, smithRoom, anvilRoom}, nil
}

func (sb SmithBuilding) getOwner() player.Player {
	return sb.Owner
}
