package buildings


type Point struct {
	X int
	Y int
}

type Room struct {
	Map []Point
	WallColor Color
}

type boundedRect struct {
	minX int
	maxX int
	minY int
	maxY int
}

func GetPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (b boundedRect) isOutsideRect(p Point) bool {
	return p.X < b.minX || p.X > b.maxX || p.Y < b.minY || p.Y > b.maxY
}

func (r Room) intersectWithRaycast(p Point, i, j, vx, vy int) bool {
	riY := r.Map[i].Y + vy
	rjY := r.Map[j].Y + vy
	riX := r.Map[i].X + vx
	rjX := r.Map[j].X + vx

	firstCondition := ( (riY > p.Y) != (rjY > p.Y) )
	if !firstCondition {
		return false
	}
	secondCondition := (p.X < (rjX - riX) * (p.Y - riY) / (rjY - riY) + riX)
	return firstCondition && secondCondition
}


func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
