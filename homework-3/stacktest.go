package main

import (
	"GoCourse/homework-3/stack"
	"fmt"
)

func main() {
	stack.Push("Один")
	stack.Push("Два")
	stack.Push("Три")

	for {
		if el, err := stack.Pop(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println( el )
		}
	}
}
