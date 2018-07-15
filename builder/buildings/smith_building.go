package buildings

import (
	"github.com/Mowinski/go-design-patterns/builder/player"
)

type SmithBuilding struct {
	Owner player.Player
	X int
	Y int
}

func (sb SmithBuilding) Build() (Building, error) {
	var building Building
	var err error
	building.Rooms, err = sb.BuildRooms()
	building.Endurance = 900.00
	building.Owner = sb.GetOwner()
	building.X = sb.X
	building.Y = sb.Y
	
	return building, err
}

func (hb SmithBuilding) BuildRooms() ([]Room, error) {
	mainRoom := Room{
		Map: []Point{GetPoint(0, 0), GetPoint(10, 0), GetPoint(10, 10), GetPoint(0, 10)},
		WallColor: YELLOW,
		RoomLetter: 'm',
	}
	
	smithRoom := Room{
		Map: []Point{GetPoint(10, 0), GetPoint(20, 0), GetPoint(20, 5), GetPoint(10, 5)},
		WallColor: BLUE,
		RoomLetter: 's',
	}

	anvilRoom := Room{
		Map: []Point{GetPoint(0, 10), GetPoint(5, 10), GetPoint(5, 15), GetPoint(0, 15)},
		WallColor: BLUE,
		RoomLetter: 'a',
	}
	
	return []Room{mainRoom, smithRoom, anvilRoom}, nil
}

func (sb SmithBuilding) GetOwner() player.Player {
	return sb.Owner
}
