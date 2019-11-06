package basecars

type PassengerCar struct {
	BaseCar
	passCapacity float64
	passFill float64
}

func NewPassengerCar() PassengerCar {  
	return PassengerCar{}
}

func (pc PassengerCar) Capacity() float64 {
        return pc.passCapacity
}

func ( pc *PassengerCar ) SetCapacity(passCapacity float64) {
        pc.passCapacity = passCapacity
}
