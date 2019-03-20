package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	//对2取 膜  / 2
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(4545645645665))
}
