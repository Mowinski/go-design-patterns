package buildings

import (
	"../player"
)

type ShopBuilding struct {
	Owner player.Player
	X int
	Y int
}


func (sb ShopBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.BuildRooms()
	building.Endurance = 700.00
	building.Owner = sb.GetOwner()
	building.X = sb.X
	building.Y = sb.Y
	return building, err
}


func (hb ShopBuilding) BuildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,
		RoomLetter: 'm',
	}
	
	storeRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(40, 0), GetPoint(40, 5), GetPoint(10, 5)},
		WallColor: BLUE,
		RoomLetter: 's',
	}

	toiletRoom := Room{
		Map: []Point{GetPoint(10, 5), GetPoint(12, 5), GetPoint(12, 7), GetPoint(10, 7)},
		WallColor: BLUE,
		RoomLetter: 't',
	}
	
	managementRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(4, 10), GetPoint(4, 14), GetPoint(0, 14)},
		WallColor: RED,
		RoomLetter: 'c',
	}
	
	return []Room{mainRoom, storeRoom, toiletRoom, managementRoom}, nil
}

func (sb ShopBuilding) GetOwner() player.Player {
	return sb.Owner
}
