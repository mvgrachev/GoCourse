package main

import (
	"fmt"
	"GoCourse/homework-3/basecars"
)

func main() {
	var passengerCar = basecars.PassengerCar{BaseCar: basecars.BaseCar{Brand: "Opel", OpenWindows: false, YearManufacture: 2012, StartEngine: true}, PassCapacity: 1100}
	var truckCar = basecars.TruckCar{BaseCar: basecars.BaseCar{Brand: "Man", OpenWindows: false, YearManufacture: 1995, StartEngine: false}, TruckCapacity: 100500}
	fmt.Println(passengerCar.PassCapacity)
	fmt.Println(truckCar.TruckCapacity)
	passengerCar.PassCapacity = 900
	truckCar.TruckCapacity = 50500
	fmt.Println(passengerCar.PassCapacity)
	fmt.Println(truckCar.TruckCapacity)
	fmt.Println(passengerCar)
	fmt.Println(truckCar)
}
