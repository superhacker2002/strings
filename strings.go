package strings

func ReverseString(str string) string {
	byteString := []rune(str)
	for i, j := 0, len(byteString)-1; i < j; i, j = i+1, j-1 {
		byteString[i], byteString[j] = byteString[j], byteString[i]
	}
	return string(byteString)
}

func CharactersCount(str string) int {
	var n int
	for range str {
		n++
	}
	return n
}
