package question

import "fmt"

func QuestionFour(str1, str2 string) {
	fmt.Println(twoString(str1, str2))
}
func twoString(str1 string, str2 string) (answer string) {

	answer = ""
	str1 += "~"
	str2 += "~"
	i := 0
	j := 0

	runeS1 := []rune(str1)
	runeS2 := []rune(str2)

	for string(runeS1[i]) != "" || string(runeS2[j]) != "" {
		if string(runeS1[i]) != "~" && str1[i:] < str2[j:] {
			answer += string(runeS1[i])
			i += 1
		} else {
			if string(runeS2[j]) == "~" {
				break
			}
			answer += string(runeS2[j])
			j += 1
		}
	}

	return
}
