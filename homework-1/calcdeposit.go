package main

import "GoCourse/homework-1/deposit"

func main() {
    var sumRur, percent float64
    deposit.GetDepositParams(&sumRur, &percent)
    deposit.CalcDeposit(sumRur,percent)
}
