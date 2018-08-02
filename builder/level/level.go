package level

import (
	"fmt"
	"image"
	"image/draw"
	"image/color"
	"image/jpeg"
	"os"

	"../buildings"
)

type Level struct {
	SizeX int
	SizeY int

	Buildings []buildings.Building
}

type Representation [][]rune

func NewLevel(sizeX, sizeY int, buildings []buildings.Building) Level {
	var level Level
	level.SizeX = sizeX
	level.SizeY = sizeY
	level.Buildings = buildings

	return level
}

func (l Level) Render() Representation {
	level :=  make([][]rune, l.SizeX)
	for x := 0; x < l.SizeX; x++ {
		level[x] = make([]rune, l.SizeY)
		for y := range level[x] {
			level[x][y] = l.getLetter(x, y)
		}
	}
	return level
}


func (r Representation) ShowInConsole() {
	for _, row := range r {
		for _, value := range row {
			fmt.Printf("%c", value)
		}
		fmt.Println()
	}
}
func (r Representation) SaveToFile(fileName string, sizeX, sizeY int) error {
	out, err := os.Create(fileName)
	defer out.Close()

	if err != nil {
		return err
	}

	var opt jpeg.Options
	img := r.drawRGBAImage(sizeX, sizeY)
	opt.Quality = 100
	err = jpeg.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r Representation) drawRGBAImage(sizeX, sizeY int) *image.RGBA {
	imgRect := image.Rect(0, 0, sizeX, sizeY)
	img := image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			draw.Draw(img, image.Rect(x, y, x+1, y+1), r.getColor(x, y), image.ZP, draw.Src)
		}
	}

	return img
}

func (r Representation) getColor(x, y int) *image.Uniform {
	var retColor color.RGBA

	switch r[y][x] {
		case 'm':
			retColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
		case 's':
			retColor = color.RGBA{R: 255, G: 255, B: 0, A: 0}
		case 't':
			retColor = color.RGBA{R: 255, G: 0, B: 0, A: 0}
		case 'c':
			retColor = color.RGBA{R: 0, G: 0, B: 255, A: 0}
		case 'a':
			retColor = color.RGBA{R: 0, G: 255, B: 255, A: 0}
		case 'o':
			retColor = color.RGBA{R: 255, G: 0, B: 255, A: 0}
		default:
			retColor = color.RGBA{R: 255, G: 255, B: 255, A: 0}
	}
	return &image.Uniform{C: retColor}
}

func (l Level) getLetter(x, y int) rune {
	letter := ' '
	point := buildings.GetPoint(x, y)
	for _, build := range l.Buildings {
		for _, room := range build.Rooms {
			if build.IsInRoom(point, room) {
				return room.RoomLetter
			}
		}
	}
	return letter
}
