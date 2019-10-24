package curconv

import "fmt"

//SayUsd выведет на экран сумму в долларах
func SayUsd(rur float64) {
    const usd float64 = 63.69
    var total float64 = rur/usd
    fmt.Printf("Сумма в долларах: %.02f\n", total)
}
