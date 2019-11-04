package basecars

type TruckCar struct {
	BaseCar
	truckCapacity float64
	truckFill float64
}

func NewTruckCar() TruckCar {
        return TruckCar{}
}

func (tc TruckCar) Capacity() float64 {
        return tc.truckCapacity
}

func ( tc *TruckCar ) SetCapacity(truckCapacity float64) {
        tc.truckCapacity = truckCapacity
}
