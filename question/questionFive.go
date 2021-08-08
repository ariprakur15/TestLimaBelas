package question

func QuestionFive(n int32,queries [][]int32) int64 {
	return arrayManipulation(n,queries)
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)

	for _, element := range queries {
		a := int(element[0])
		b := element[1]
		k := int64(element[2])
		arr[a] += k
		if (b + 1) <= n {
			arr[b+1] -= k
		}
	}

	var x, max int64 = 0, 0
	for i := 1; i <= int(n); i++ {
		x += arr[i]
		if max < x {
			max = x
		}
	}
	return max
}
