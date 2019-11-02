package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	n := enterNumber("Введите число:")
	checkNum(n)
}

func checkNum( n int32 ) {
	if n%2 == 0 {
		fmt.Println("Число четное")
	} else if n%3 == 0 {
		fmt.Println("Число делится без остатка на 3")
	} else {
		fmt.Println("Число нечетное")
	}
}

func inputData(msg string) (data string){
    fmt.Println(msg)
    fmt.Scanln(&data)
    if data == "exit" {
        panic(nil)
    }
    return
}

func enterNumber(msg string) int32 {
	for {
		num, err := strconv.ParseInt(inputData(msg), 10, 32)
		if err != nil {
			fmt.Println("Не удалось распознать число")
		} else {
			// Ввод числа прошел без ошибок
			return int32(num)
			break
		}
	}
	return 0
}
