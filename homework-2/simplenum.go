package main

import "fmt"

func main() {
	result := getSimpleNum()

	fmt.Println(cap(result));
	fmt.Println(result)
}

func getSimpleNum() ([]int) {
	arraySimpleNum := make([]int,100)
	p := 2
	for i := 0; i < 100; i++ {
		arraySimpleNum[i] = p
		p++
	}

	isComplete := 0
	for {
		isComplete = 1	
		for ind := 0; ind < len(arraySimpleNum); ind++ {
			for index, v := range arraySimpleNum {
				if ( v > arraySimpleNum[ind] ) {
					val := arraySimpleNum[ind]
					if v%val == 0 {
						lastCurEl := arraySimpleNum[len(arraySimpleNum)-1]+1
						splice(&arraySimpleNum,index,1)
						arraySimpleNum = append(arraySimpleNum,lastCurEl)
						isComplete = 0
					}
				}
			}
		}
		if isComplete == 1 {
			break
		}
	}

	return arraySimpleNum
}

func splice ( s *[]int, offset int, limit int ) {
	*s = append( (*s)[:offset], (*s)[offset+limit:]...)
}
