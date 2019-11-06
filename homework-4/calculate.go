package main

import (
	"GoCourse/homework-4/calculator"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			fmt.Println("Справка по использованию калькулятора:")
			fmt.Printf("%6s - %6s\n", "help", "выведет справку")
			fmt.Printf("%6s - %6s\n", "exit", "завершит работу калькулятора")
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
