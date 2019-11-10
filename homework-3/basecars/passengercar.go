package basecars

import "fmt"
type PassengerCar struct {
	BaseCar
	passCapacity float64
	passFill float64
}

func NewPassengerCar(brand string) PassengerCar { 
	bc := BaseCar{ 
		brand: brand,
		yearManufacture: 0,
		stateEngine: "start",
		stateWindows: "open",
		power: 100,
	}
	return PassengerCar{ bc, 0, 0 }
}

func (pc PassengerCar) Capacity() float64 {
        return pc.passCapacity
}

func ( pc *PassengerCar ) SetCapacity(passCapacity float64) {
        pc.passCapacity = passCapacity
}
