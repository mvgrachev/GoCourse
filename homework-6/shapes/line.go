package shapes

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)
 
type Line struct {
	startPoint image.Point
	endPoint image.Point
	img image.RGBA
}

//Сконструируем нашу линию
func NewLine(x1, y1, x2, y2 int) *Line{
	startPoint := image.Point{X:x1,Y:y1}
	endPoint := image.Point{X:x2,Y:y2}
	
	col := color.RGBA{255, 0, 0, 255}
	lineImg := image.NewRGBA(image.Rect(x1, y1, x2, y2))
        draw.Draw(lineImg, lineImg.Bounds(), &image.Uniform{color.RGBA{0, 0, 0, 0}}, image.ZP, draw.Src)
	// здесь нужен адекватный алгритм отрисовки линии. Сейчас алгоритм неправильнышй
	// с каждым увеличением Х будем увеличивать Y на 1
	// получим наклонную линию
	stepX := x2-x1
	stepY := y2-y1
	if stepX >= stepY {
		for ; x1 <= x2; x1++ {
			lineImg.Set(x1, y1, col)
			if y2 > y1 {
				y1++
			}
		}
	} else {
		for ; y1 <= y2; y1++ {
			lineImg.Set(x1, y1, col)
			if x2 > x1 {
				x1++
			}
		}
	}
	
	return &Line{ startPoint, endPoint, *lineImg }
}

//Метод записывет изображение в файл
func (l Line) SaveToFile(pathFile string) {
        file, err := os.Create(pathFile)
        if err != nil {
                log.Fatalf("Failed create file: %s", err)
        }
        defer file.Close()
        png.Encode(file, &l.img)
}
