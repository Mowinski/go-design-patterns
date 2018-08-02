package buildings

import (
	"../player"
)


type Building struct {
	Rooms []Room
	Endurance float32
	Owner player.Player
	X int
	Y int
}

func (b Building) IsInRoom(p Point, r Room) bool {
	boundedRect := b.getBoundedRect(r)

    if boundedRect.isOutsideRect(p) {
        return false
    }
    if isPointOnLine(p, r) {
    	return true
	}
	inside := false
	j := len(r.Map) - 1
    for i := 0; i < len(r.Map); i++ {
		if r.intersectWithRaycast(p, i, j, b.X, b.Y) {
           inside = !inside
		}
		j = i
    }

    return inside
}

func isPointOnLine(point Point, room Room) bool {
	// TODO implement it

	return true
}


func (b Building) getBoundedRect(r Room) boundedRect {
	rect := boundedRect{
		minX: r.Map[0].X + b.X,
		maxX: r.Map[0].X + b.X,
		minY: r.Map[0].Y + b.Y,
		maxY: r.Map[0].Y + b.Y,
	}
    for i := 1; i < len(r.Map); i++ {
		q := r.Map[i] 
		tx := q.X + b.X
		ty := q.Y + b.Y
        rect.minX = min(tx, rect.minX)
        rect.maxX = max(tx, rect.maxX)
        rect.minY = min(ty, rect.minY)
        rect.maxY = max(ty, rect.maxY)
	}
	
	return rect
}
