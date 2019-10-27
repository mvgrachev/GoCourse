package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Println("Ввод количества чисел фибоначи:")
	fmt.Scanln(&n)

	a := big.NewInt(0)
	b := big.NewInt(1)

	fmt.Println("Выводим список чисел фибоначи:")
	fmt.Println(a)

	for i := 1; i < n; i++ {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)	
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a

		fmt.Println(a)
	}
}
