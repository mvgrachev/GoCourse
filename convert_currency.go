package main

import "fmt"

func main() {
    const usd float64 = 63.69
    var rur float64
    fmt.Println("Введите сумму в рублях")
    fmt.Scanln(&rur)
    var total float64 = rur/usd
    fmt.Printf("Сумма в долларах: %.02f\n", total)
}
