package deposit

import (
	"fmt"
)

//CalcDeposit выведет на экран итоговую сумму вклада (простой процент)
func CalcDeposit(sumRur float64, percent float64) {
    const year float64 = 5
    var diff float64 = ( sumRur * percent )/100
    var totalSum float64 = sumRur + diff * year
    fmt.Printf("Итоговая сумма через 5 лет: %.02f\n", totalSum)
}
