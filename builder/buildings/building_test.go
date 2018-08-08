package buildings_test

import (
	"testing"

	"../buildings"
)

func TestInSquareRoomAt0x0(t *testing.T) {
	testRoom := buildings.Room{
		Map: []buildings.Point{buildings.GetPoint(2, 2), buildings.GetPoint(2, 4), buildings.GetPoint(4, 4), buildings.GetPoint(4, 2)},
	}

	testBuilding := buildings.Building{Rooms: []buildings.Room{testRoom}, X: 0, Y: 0}
	
	result := testBuilding.IsInRoom(buildings.GetPoint(3, 3), testRoom)

	if result != true {
		t.Error("Point (3, 3) should be in buildings. Test fails")
	}

	result = testBuilding.IsInRoom(buildings.GetPoint(1, 1), testRoom) 
	
	if result != false {
		t.Error("Point (1, 1) shouldn't be in buildings. Test fails")
	}
}


func TestInSquareRoomAtMovedBuilding(t *testing.T) {
	testRoom := buildings.Room{
		Map: []buildings.Point{buildings.GetPoint(2, 2), buildings.GetPoint(2, 4), buildings.GetPoint(4, 4), buildings.GetPoint(4, 2)},
	}

	testBuilding := buildings.Building{Rooms: []buildings.Room{testRoom}, X: 2, Y: 2}
	
	result := testBuilding.IsInRoom(buildings.GetPoint(5, 5), testRoom)

	if result != true {
		t.Error("Point (5, 5) should be in buildings. Test fails")
	}

	result = testBuilding.IsInRoom(buildings.GetPoint(1, 1), testRoom) 
	
	if result != false {
		t.Error("Point (1, 1) shouldn't be in buildings. Test fails")
	}
}
