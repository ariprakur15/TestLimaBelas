package main

import (
	"LimaBelasInnovate/question"
	"fmt"
)

func main() {
	fmt.Println("#1 Result")
	question.QuestionOne(3, 5)
	fmt.Println("#2 Result")
	question.QuestionTwo([]int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6})
	fmt.Println("#3 Result")
	question.QuestionThree("saya", "kamu")
	fmt.Println("#4 Result")
	question.QuestionFour("jack", "daniel")
	fmt.Println("#5 Result")
	question.QuestionFive(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}})
}
