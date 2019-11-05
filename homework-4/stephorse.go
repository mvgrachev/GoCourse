package main

import (
	"fmt"
	"errors"
)

func main() {
	var positionSymbol string
	var positionNumber int
	getCurrentPosition(&positionSymbol, &positionNumber)
	currentPositionHorse := positionHorse{ Symbol: positionSymbol, Number: positionNumber}
	fmt.Println(currentPositionHorse)
	if success, err := currentPositionHorse.checkPosition(); !success {
		fmt.Println(err)
	}
	
	enablePositionsHorse := currentPositionHorse.calcEnablePositions()
	fmt.Println(enablePositionsHorse)
	for _,enablePos := range enablePositionsHorse {
		enablePosChilds := enablePos.calcEnablePositions()
		fmt.Println(enablePosChilds)
	}
}

type positionHorse struct {
	Symbol string
	Number int
}

var allowSymbols = map[string]int{"a":0,"b":1,"c":2,"d":3,"e":4,"f":5,"g":6,"h":7}
var allowNumbers = map[int]int{1:0,2:1,3:2,4:3,5:4,6:5,7:6,8:7}

func getCurrentPosition( positionSymbol *string, positionNumber *int ) {
	fmt.Println("Введите букву")
	fmt.Scanln(positionSymbol)
	fmt.Println("Введите цифрру")
	fmt.Scanln(positionNumber)
}

func ( pH positionHorse) checkPosition() ( success bool, err error ) {
	isError := false
	if _, ok := allowSymbols[pH.Symbol]; !ok {    
		isError = true
	}
	if _, ok := allowNumbers[pH.Number]; !ok {
		isError = true    
	}

	if isError {	
		err = errors.New("Позиция выходит за пределы шахматной доски")
	} else {
		success = true
	}
	return
}

func (pH positionHorse) calcEnablePositions() ( enablePositionsHorse []positionHorse ) {
	var Symbols = [8]string{"a","b","c","d","e","f","g","h"}
	var Numbers = [8]int{1,2,3,4,5,6,7,8}
	// В переменной содержатся допустимые переходы относительно текущей позиции
	var allowMoves = [8][2]int{{-2,-1},{-2,1},{2,-1},{2,1},{-1,-2},{-1,2},{1,-2},{1,2}}
	for _, values := range allowMoves {
		if allowSymbols[pH.Symbol] - values[0] < 0 {
			continue
		}
		if allowNumbers[pH.Number] - values[1] < 0 {
			continue
		}
		newPositionHorse := positionHorse{Symbol:Symbols[allowSymbols[pH.Symbol] - values[0]], Number:Numbers[allowNumbers[pH.Number] - values[1]]}
		enablePositionsHorse = append(enablePositionsHorse,newPositionHorse)
	}

	return
}
