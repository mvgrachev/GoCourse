package main

import (
	"fmt"
	"GoCourse/homework-3/basecars"
)

func main() {
	var passengerCar = basecars.PassengerCar{BaseCar: basecars.BaseCar{Brand: "Opel", OpenWindows: false, YearManufacture: 2012, StateEngine: "started"}, PassCapacity: 1100}
	var aPC basecars.AllCar
	aPC = &passengerCar
	var truckCar = basecars.TruckCar{BaseCar: basecars.BaseCar{Brand: "Man", OpenWindows: false, YearManufacture: 1995, StateEngine: "stoped"}, TruckCapacity: 100500}
	var aTC basecars.AllCar
	aTC = &truckCar
	fmt.Println(aPC)
	fmt.Println(aTC)
	aPC.StopEngine()
	aTC.StartEngine()
	fmt.Println(passengerCar.PassCapacity)
	fmt.Println(truckCar.TruckCapacity)
	passengerCar.PassCapacity = 900
	truckCar.TruckCapacity = 50500
	fmt.Println(passengerCar.PassCapacity)
	fmt.Println(truckCar.TruckCapacity)
	fmt.Println(passengerCar)
	fmt.Println(truckCar)
}
