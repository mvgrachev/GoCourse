package main

import (
	"fmt"
	"GoCourse/homework-3/basecars"
)

func totalCapacity(cars ...basecars.AllCar) float64 {
	var capacity float64
	for _, c := range cars {
		if c.Capacity != nil {
			capacity += c.Capacity()
		}
	}

	return capacity
}

func main() {
	passengerCar := basecars.NewPassengerCar()
	passengerCar.SetBrand("Opel")	
	passengerCar.SetCapacity(1000)	
	
	passengerCar1 := basecars.NewPassengerCar()
	passengerCar1.SetBrand("Mazda")	
	passengerCar1.SetCapacity(600)	
	
	truckCar := basecars.NewTruckCar()
	truckCar.SetBrand("MAN")	
	truckCar.SetCapacity(100000)	
	
	truckCar1 := basecars.NewTruckCar()
	truckCar1.SetBrand("Volvo")	
	truckCar1.SetCapacity(100000)	
	
	fmt.Println(passengerCar)
	fmt.Println(passengerCar1)
	fmt.Println(truckCar)
	fmt.Println(truckCar1)

	fmt.Println("Объем багажников легковых авто: ", totalCapacity(&passengerCar,&passengerCar1))
	fmt.Println("Объем кузовов грузовых авто: ", totalCapacity(&truckCar,&truckCar1))
	fmt.Println("Общий объем багажников всех авто: ", totalCapacity(&passengerCar,&passengerCar1,&truckCar,&truckCar1))
}
