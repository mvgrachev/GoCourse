package main

import "GoCourse/deposit"

func main() {
    var sumRur float64
    var percent float64
    deposit.GetDepositParams(&sumRur, &percent)
    deposit.CalcDeposit(sumRur,percent)
}
