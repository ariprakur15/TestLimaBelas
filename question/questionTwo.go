package question

import (
	"fmt"
)

func QuestionTwo(arrData []int) {
	getValue(arrData)
}

func getValue(arrData []int) {
	var getValuMap = make(map[int]int)
	var arrSort  []int
	for _, val := range arrData {
		if getValuMap[val] == 0 {
			arrSort = append(arrSort,val)
		}
		getValuMap[val] += 1
	}

	if len(arrSort) != 0 {
		fmt.Println("number --> total")
		for _, val := range arrSort {
			fmt.Println(val , " --> ", getValuMap[val])
		}
	}
}
