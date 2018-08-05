package actors_test

import (
	"testing"
	"../actors"
)

func assertTrue(t *testing.T, result bool, element string) {
	if result != true {
		t.Errorf("Expected true, got false in element %s", element)
	}
}

func assertFalse(t *testing.T, result bool, element string) {
	if result != false {
		t.Errorf("Expected false, got true in element %s", element)
	}
}


func TestIsValidDirection(t *testing.T) {
	assertTrue(t, actors.IsValidDirection('a'), "a")
	assertTrue(t, actors.IsValidDirection('a'), "s")
	assertTrue(t, actors.IsValidDirection('a'), "d")

	assertFalse(t, actors.IsValidDirection('f'), "f")
	assertFalse(t, actors.IsValidDirection('q'), "q")
	assertFalse(t, actors.IsValidDirection('w'), "w")
	assertFalse(t, actors.IsValidDirection('e'), "e")
	assertFalse(t, actors.IsValidDirection('z'), "z")
	assertFalse(t, actors.IsValidDirection('x'), "x")
	assertFalse(t, actors.IsValidDirection('c'), "c")
}
