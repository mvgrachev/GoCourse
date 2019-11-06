package basecars

type BaseCar struct {
	brand string
	yearManufacture int
	stateEngine string
	stateWindows string
	power float64
}

func (c BaseCar) Brand() string {
	return c.brand
}

func (c *BaseCar) SetBrand( brand string){
	c.brand = brand
}

func (c BaseCar) YearManufacture() int {
	return c.yearManufacture
}

func (c *BaseCar) SetYearManufacture(yearManufacture int) {
	c.yearManufacture = yearManufacture
}

func (c BaseCar) StateEngine() string {
	return c.stateEngine
}

func ( c *BaseCar ) SetStateEngine(stateEngine string) {
	c.stateEngine = stateEngine
}

func (c BaseCar) StateWindows() string {
	return c.stateWindows
}

func ( c *BaseCar ) SetStateWindows(stateWindows string) {
	c.stateWindows = stateWindows
}

func (c BaseCar) Power() float64 {
	return c.power
}

func ( c *BaseCar ) SetPower(power float64) {
	c.power = power
}
