package question

import "fmt"

func QuestionThree(str1, str2 string) {
	fmt.Println(mergeString(str1, str2))
}
func mergeString(string1 string, string2 string) (result string) {

	runeS1 := []rune(string1)
	runeS2 := []rune(string2)

	for i := 0; i < len(string1) || i < len(string2); i++ {

		if i < len(string1) {
			result += string(runeS1[i])
		}

		if i < len(string2) {
			result += string(runeS2[i])

		}

	}

	return

}
