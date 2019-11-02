package basecars

type BaseCar struct {
	Brand string
	YearManufacture int
	StateEngine string
	OpenWindows bool
}

func ( c *BaseCar ) StartEngine() {
	c.StateEngine = "started"
}

func ( c *BaseCar ) StopEngine() {
	c.StateEngine = "stoped"
}
