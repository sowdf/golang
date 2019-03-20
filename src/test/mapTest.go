package main

import "fmt"

func fundMaxLength(s string) int {
	start := 0
	maxLength := 0
	lastCurrentCode := make(map[byte]int)
	for i, ch := range []byte(s) {
		fmt.Println(ch)
		if lastI, ok := lastCurrentCode[ch]; ok && lastI >= start {
			start = lastCurrentCode[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastCurrentCode[ch] = i
	}
	return maxLength
}

func main() {
	a := "abcdefg"
	fmt.Println([]byte(a))
	fmt.Println(fundMaxLength("aabcbd"))
}
