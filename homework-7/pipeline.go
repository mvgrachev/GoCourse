package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func(naturals chan int) {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}(naturals)

	// возведение в квадрат
	go func(naturals chan int,squares chan int) {
		for {
			x, ok := <-naturals
			
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}(naturals,squares)

	// печать
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
