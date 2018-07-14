package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type SmithBuilding struct {
	owner player.Player
}

func (sb SmithBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.BuildRoom()
	building.Endurance = 900.00
	building.Owner = sb.GetOwner()
	
	return building, err
}

func (hb SmithBuilding) BuildRoom() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,		
	}
	
	smithRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(20, 0), GetPoint(20, 5), GetPoint(10, 5)},
		WallColor: BLUE,
	}

	anvilRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(5, 10), GetPoint(5, 15), GetPoint(0, 15)},
		WallColor: BLUE,
	}
	
	return []Room{mainRoom, smithRoom, anvilRoom}, nil
}

func (sb SmithBuilding) GetOwner() player.Player {
	return sb.owner
}
