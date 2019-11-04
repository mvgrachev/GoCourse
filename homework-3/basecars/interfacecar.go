package basecars

type AllCar interface {
	Brand() string
	SetBrand(brand string)
	YearManufacture() int
	SetYearManufacture(yerManufacture int)
	StateEngine() string
	SetStateEngine(stateEngine string)
	StateWindows() string
	SetStateWindows(stateWindows string)
	Power() float64
	SetPower(power float64)
	Capacity() float64
	SetCapacity(power float64)
} 
