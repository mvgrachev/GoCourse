package main

import "GoCourse/curconv"

func main() {
    var rur float64
    curconv.GetRur(&rur)
    curconv.SayUsd(rur)
}
