package deposit

import "fmt"

// GetDepositParams попросит ввести сумму вклада в рублях и годовой процент
func GetDepositParams( sumRur *float64, percent *float64 ) {
    fmt.Println("Введите сумму в рублях")
    fmt.Scanln(sumRur)
    fmt.Println("Введите годовой процент")
    fmt.Scanln(percent)
}
