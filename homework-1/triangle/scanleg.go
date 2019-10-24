package triangle

import "fmt"

// GetLegs попросит ввести катеты треугольника
func GetLegs( a *float64, b *float64 ) {
    fmt.Println("Введите первый катет")
    fmt.Scanln(a)
    fmt.Println("Введите второй катет")
    fmt.Scanln(b)
}
