package main

import "GoCourse/homework-1/curconv"

func main() {
    var rur float64
    curconv.GetRur(&rur)
    curconv.SayUsd(rur)
}
