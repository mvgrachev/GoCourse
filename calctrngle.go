package main

import "GoCourse/triangle"

func main() {
    var a float64
    var b float64
    triangle.GetLegs(&a, &b)
    triangle.CalcTriangle(a,b)
}
