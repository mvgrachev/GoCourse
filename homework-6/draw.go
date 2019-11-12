package main

import (
	"GoCourse/homework-6/shapes"
)

func main() {
	shapes.DrawRectangle()
	line := shapes.NewLine(0,0,100,100)
	line.SaveToFile("first_line.png")
}
