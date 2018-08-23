package buildings

import (
	"../player"
)

type ShopBuilding struct {
	IBuilding

	Owner player.Player
	X int
	Y int
}


func (sb ShopBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.buildRooms()
	building.Endurance = 700.00
	building.Owner = sb.getOwner()
	building.X = sb.X
	building.Y = sb.Y
	return building, err
}


func (sb ShopBuilding) buildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		RoomLetter: 'm',
	}
	
	storeRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(40, 0), GetPoint(40, 5), GetPoint(10, 5)},
		RoomLetter: 's',
	}

	toiletRoom := Room{
		Map: []Point{GetPoint(10, 5), GetPoint(12, 5), GetPoint(12, 7), GetPoint(10, 7)},
		RoomLetter: 't',
	}
	
	managementRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(4, 10), GetPoint(4, 14), GetPoint(0, 14)},
		RoomLetter: 'c',
	}
	
	return []Room{mainRoom, storeRoom, toiletRoom, managementRoom}, nil
}

func (sb ShopBuilding) getOwner() player.Player {
	return sb.Owner
}
