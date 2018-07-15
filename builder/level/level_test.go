package level_test

import (
	"testing"

	"github.com/Mowinski/go-design-patterns/builder/level"
	"github.com/Mowinski/go-design-patterns/builder/buildings"
)


func TestLevelRenderSize(t *testing.T) {
	var buildings []buildings.Building
	testLevel := level.NewLevel(10, 10, buildings)
	levelCharacter := testLevel.Render()
	
	if len(levelCharacter) != 10 {
		t.Errorf("Level has wrong X size, expected 10, got %d", len(levelCharacter))
	}

	if len(levelCharacter[0]) != 10 {
		t.Errorf("Level has wrong Y size, expected 10, got %d", len(levelCharacter[0]))
	}
}

func TestEmptyLevel5x5(t *testing.T) {
	var buildings []buildings.Building
	testLevel := level.NewLevel(5, 5, buildings)
	levelCharacter := testLevel.Render()

	for y, row := range levelCharacter {
		for x, point := range row {
			if point != ' ' {
				t.Errorf("At %dx%d element is not empty", x, y)
			}
		}
	}
}

func TestOneBuildingWithOneRoomLevel5x5(t *testing.T) {
	testRoom := buildings.Room{
		Map: []buildings.Point{
			buildings.GetPoint(0, 0),
			buildings.GetPoint(0, 1),
			buildings.GetPoint(1, 1),
			buildings.GetPoint(1, 0),
			buildings.GetPoint(0, 0),
		},
		WallColor: buildings.RED,
		RoomLetter: 'b',
	}
	testBuilding := buildings.Building{
		Rooms: []buildings.Room{testRoom},
		X: 3,
		Y: 3,
	}
	buildings := []buildings.Building{testBuilding}
	testLevel := level.NewLevel(5, 5, buildings)
	levelCharacter := testLevel.Render()
	
	if levelCharacter[3][3] != 'b' {
		t.Errorf("At (3, 3) there is no building, expected 'b' got %c", levelCharacter[3][3])
	}

	if levelCharacter[3][4] != 'b' {
		t.Errorf("At (3, 4) there is no building, expected 'b' got %c", levelCharacter[3][4])
	}

	if levelCharacter[4][4] != 'b' {
		t.Errorf("At (4, 3) there is no building, expected 'b' got %c", levelCharacter[4][3])
	}

	if levelCharacter[4][4] != 'b' {
		t.Errorf("At (4, 4) there is no building, expected 'b' got %c", levelCharacter[4][4])
	}
}
