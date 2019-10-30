package main

import (
	"GoCourse/homework-3/stack"
	"fmt"
)

func main() {
	stack.Push("Один")
	stack.Push("Два")
	stack.Push("Три")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("Четыре")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
