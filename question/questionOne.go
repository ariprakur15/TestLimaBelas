package question

import "fmt"

func QuestionOne(a, b int) {
	fmt.Println("Before Swap:")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a, b = changeVariable(a, b)
	fmt.Println("After Swap:")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}

func changeVariable(a int, b int) (int, int) {
	b = a + b
	a = b - a
	b = b - a
	return a, b
}
