package curconv

import "fmt"

// GetRur попросит ввести сумму в рублях
func GetRur( rur *float64 ) {
    fmt.Println("Введите сумму в рублях")
    fmt.Scanln(rur)
}
