package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type ShopBuilding struct {
	owner player.Player
}


func (sb ShopBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.BuildRoom()
	building.Endurance = 700.00
	building.Owner = sb.GetOwner() 
	return building, err
}


func (hb ShopBuilding) BuildRoom() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,		
	}
	
	storeRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(40, 0), GetPoint(40, 5), GetPoint(10, 5)},
		WallColor: BLUE,
	}

	toiletRoom := Room{
		Map: []Point{GetPoint(10, 5), GetPoint(12, 5), GetPoint(12, 7), GetPoint(10, 7)},
		WallColor: BLUE,
	}
	
	managementRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(4, 10), GetPoint(4, 14), GetPoint(0, 14)},
		WallColor: RED,
	}
	
	return []Room{mainRoom, storeRoom, toiletRoom, managementRoom}, nil
}

func (sb ShopBuilding) GetOwner() player.Player {
	return sb.owner
}
