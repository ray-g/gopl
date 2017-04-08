package reverse

func reverse(arr *[10]int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}
