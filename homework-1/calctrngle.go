package main

import "GoCourse/homework-1/triangle"

func main() {
    var a, b float64
    triangle.GetLegs(&a, &b)
    triangle.CalcTriangle(a,b)
}
