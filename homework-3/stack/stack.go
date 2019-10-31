package stack

import "errors"

var x []string

// Push добавит новый элемент в стек
func Push(str string) {
        x = append(x, str)
}

// Pop вернет последний добавленный в стек элемент
func Pop() (popElem string, err error) {
        numOfElements := len(x)
        // Когда стек будет пустым, он вернет пустую строку
        if numOfElements == 0 {
                err = errors.New("Stack is empty")
		return
        }
        popElem = x[numOfElements-1]
        x = x[:numOfElements-1]

        return
}
