package triangle

import (
	"fmt"
	"math"
)

//CalcTriangle выведет на экран гипотенузу, периметр и площадь
func CalcTriangle(a float64, b float64) {
    var hypotenuse float64 = math.Sqrt(math.Pow(a,2) + math.Pow(b,2))
    var perimeter float64 = a + b + hypotenuse
    var area float64 = (a*b)/2
    fmt.Printf("Гипотенуза: %.02f\n", hypotenuse)
    fmt.Printf("Периметр: %.02f\n", perimeter)
    fmt.Printf("Площадь: %.02f\n", area)
}
